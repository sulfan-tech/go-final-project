package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-final-project/internal/delivery/handlers/user"
	usermodel "go-final-project/internal/domain/user/model"
	"go-final-project/internal/domain/user/service/mocks"
)

func TestUserLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a mock user service
	mockUserService := new(mocks.MockUserService)
	handler := user.NewUserHandler(mockUserService)

	// Create a Gin router
	router := gin.Default()
	router.POST("/login", handler.UserLogin)

	// Mock login request
	loginRequest := user.LoginRequest{
		Email:    "test@test.com",
		Password: "password",
	}
	body, _ := json.Marshal(loginRequest)

	// Mock the user returned by the service
	expectedUser := &usermodel.User{
		ID:       1,
		Username: "test",
		Email:    "test@test.com",
		Age:      30,
	}
	mockUserService.On("UserAuthenticate", mock.Anything, loginRequest.Email, loginRequest.Password).Return(expectedUser, nil)

	// Perform the HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var responseBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Equal(t, "test", responseBody["user"].(map[string]interface{})["username"])
	assert.Equal(t, "test@test.com", responseBody["user"].(map[string]interface{})["email"])
	assert.Equal(t, float64(30), responseBody["user"].(map[string]interface{})["age"])
}

func TestUserRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a mock user service
	mockUserService := new(mocks.MockUserService)
	handler := user.NewUserHandler(mockUserService)

	// Create a Gin router
	router := gin.Default()
	router.POST("/register", handler.UserRegister)

	// Mock registration request
	registrationRequest := user.RegistrationRequest{
		Username: "test",
		Email:    "test@test.com",
		Password: "password",
		Age:      30,
	}
	body, _ := json.Marshal(registrationRequest)

	// Mock the user returned by the service
	expectedUser := &usermodel.User{
		ID:       1,
		Username: "test",
		Email:    "test@test.com",
		Age:      30,
	}
	mockUserService.On("UserRegister", mock.Anything, mock.Anything).Return(expectedUser, nil)

	// Perform the HTTP request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(body))
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	var responseBody map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.Equal(t, float64(1), responseBody["id"])
	assert.Equal(t, "test", responseBody["username"])
	assert.Equal(t, "test@test.com", responseBody["email"])
	assert.Equal(t, float64(30), responseBody["age"])
}
