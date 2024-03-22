package userrepo_test

import (
	"errors"
	"testing"

	usermodel "go-final-project/internal/domain/user/model"
	"go-final-project/internal/domain/user/repo/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	mock := new(mocks.MockUserRepository)
	user := &usermodel.User{Username: "test", Email: "test@test.com", Password: "password", Age: 30}

	// Mock successful user creation
	mock.On("CreateUser", user).Return(user, nil)

	// Call the CreateUser method
	createdUser, err := mock.CreateUser(user)

	// Check if the result matches expectations
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
}

func TestGetUserByEmail(t *testing.T) {
	mock := new(mocks.MockUserRepository)
	email := "test@test.com"
	expectedUser := &usermodel.User{ID: 1, Username: "test", Email: email, Age: 30}

	// Mock successful user retrieval by email
	mock.On("GetUserByEmail", email).Return(expectedUser, nil)

	// Call the GetUserByEmail method
	user, err := mock.GetUserByEmail(email)

	// Check if the result matches expectations
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)
}

func TestUpdateUser(t *testing.T) {
	mock := new(mocks.MockUserRepository)
	user := &usermodel.User{ID: 1, Username: "updated_test", Email: "updated_test@test.com", Password: "updated_password", Age: 40}

	// Mock successful user update
	mock.On("UpdateUser", user).Return(nil)

	// Call the UpdateUser method
	err := mock.UpdateUser(user)

	// Check if the result matches expectations
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	mock := new(mocks.MockUserRepository)
	userID := uint(1)

	// Mock successful user deletion
	mock.On("DeleteUser", userID).Return(nil)

	// Call the DeleteUser method
	err := mock.DeleteUser(userID)

	// Check if the result matches expectations
	assert.NoError(t, err)
}

func TestGetUserByID_NotFound(t *testing.T) {
	// Create a new mock instance
	mock := new(mocks.MockUserRepository)

	// Specify the user ID for which the user is not found
	userID := uint(1)

	// Mock the GetUserByID method to return nil user and an error indicating "user not found"
	mock.On("GetUserByID", userID).Return(nil, errors.New("user not found"))

	// Call the GetUserByID method from the mock
	user, err := mock.GetUserByID(userID)

	// Check if the error is as expected
	assert.Error(t, err)
	assert.Nil(t, user)
}
