package mocks

import (
	"errors"
	usermodel "go-final-project/internal/domain/user/model"
	userrepo "go-final-project/internal/domain/user/repo"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	repoUser userrepo.UserRepositoryImpl // Ubah RepoUser menjadi repoUser
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *usermodel.User) (*usermodel.User, error) {
	args := m.Called(user)
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*usermodel.User, error) {
	args := m.Called(email)
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(id uint) (*usermodel.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, errors.New("user not found")
	}
	return args.Get(0).(*usermodel.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *usermodel.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
