package routes

import (
	"go-final-project/internal/delivery/handlers/user"
	"go-final-project/internal/delivery/middlewares/jwt"
	"go-final-project/internal/domain/user/service"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	// Initialize services
	userService := service.NewInstanceUserService()

	// Initialize handlers
	userHandler := user.NewUserHandler(userService)

	// API version 1 routes
	v1 := router.Group("/api/v1")
	{
		userV1 := v1.Group("/users")
		{
			userV1.POST("/login", jwt.ValidateJWT(userHandler.UserLogin))
			userV1.POST("/register", userHandler.UserRegister)
		}
	}
}
