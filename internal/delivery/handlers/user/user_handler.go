package user

import (
	"context"
	"go-final-project/internal/delivery/middlewares/jwt"
	usermodel "go-final-project/internal/domain/user/model"
	"go-final-project/internal/domain/user/service"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserServiceImpl
}

func NewUserHandler(userService service.UserServiceImpl) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Login handles the login request.
func (uh *UserHandler) UserLogin(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := loginRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uh.userService.UserAuthenticate(c.Request.Context(), loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateJWT(user, os.Getenv("JWT_SECRET"), time.Hour*24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: failed to generate token": err.Error()})
		return
	}

	response := LoginResponse{
		Token: token,
		User:  *user,
	}

	c.JSON(http.StatusOK, response)
}

func (uh *UserHandler) UserRegister(c *gin.Context) {
	var registrationRequest RegistrationRequest
	if err := c.ShouldBindJSON(&registrationRequest); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}

	if err := registrationRequest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user := usermodel.User{
		Username: registrationRequest.Username,
		Email:    registrationRequest.Email,
		Password: registrationRequest.Password,
		Age:      registrationRequest.Age,
	}

	user.SetPassword(registrationRequest.Password)

	usr, err := uh.userService.UserRegister(context.TODO(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, usr)
}
