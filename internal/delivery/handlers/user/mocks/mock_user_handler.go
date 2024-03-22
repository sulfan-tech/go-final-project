package mocks

import (
	"context"
	"errors"
	usermodel "go-final-project/internal/domain/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockUserService struct{}

func (mus *MockUserService) UserAuthenticate(ctx context.Context, email, password string) (*usermodel.User, error) {
	if email == "test@test.com" && password == "password" {
		user := &usermodel.User{
			ID:       1,
			Username: "test",
			Email:    "test@test.com",
			Age:      30,
		}
		return user, nil
	}
	return nil, errors.New("invalid email or password")
}

func (mus *MockUserService) UserRegister(ctx context.Context, user *usermodel.User) (*usermodel.User, error) {
	user.ID = 1
	return user, nil
}

// MockUserHandler struct
type MockUserHandler struct{}

func NewMockUserHandler() *MockUserHandler {
	return &MockUserHandler{}
}

// UserLogin mocks the login functionality of UserHandler
func (h *MockUserHandler) UserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"token": "dummy_token", "user": usermodel.User{ID: 1, Username: "test", Email: "test@test.com", Age: 30}})
}

// UserRegister mocks the user registration functionality of UserHandler
func (h *MockUserHandler) UserRegister(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"id": 1, "username": "test", "email": "test@test.com", "age": 30})
}
