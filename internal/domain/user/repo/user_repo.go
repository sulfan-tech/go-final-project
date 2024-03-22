package userrepo

import (
	"fmt"
	"sync"

	"go-final-project/config/psql"
	"go-final-project/internal/delivery/middlewares/logger"
	usermodel "go-final-project/internal/domain/user/model"

	"gorm.io/gorm"
)

var (
	instance *UserRepository
	once     sync.Once
)

type UserRepository struct {
	db      *gorm.DB
	logging logger.Logger
}

type UserRepositoryImpl interface {
	CreateUser(user *usermodel.User) (*usermodel.User, error)
	GetUserByEmail(email string) (*usermodel.User, error)
	GetUserByID(id uint) (*usermodel.User, error)
	UpdateUser(user *usermodel.User) error
	DeleteUser(id uint) error
}

func NewInstanceUserRepository() (UserRepositoryImpl, error) {
	once.Do(func() {
		db, err := psql.SetupDatabase()
		if err != nil {
			panic(err)
		}
		logger := logger.NewLogger(logger.InfoLevel)

		db.AutoMigrate(&usermodel.User{})
		instance = &UserRepository{
			db:      db,
			logging: logger,
		}
	})

	return instance, nil
}

func (ur *UserRepository) CreateUser(user *usermodel.User) (*usermodel.User, error) {
	if err := ur.db.Table("users").Create(user).Error; err != nil {
		ur.logging.Error("Failed to create user: " + err.Error())
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*usermodel.User, error) {
	var user usermodel.User
	if result := ur.db.Table("users").Where("email = ?", email).Find(&user); result.Error != nil {
		ur.logging.Error("Failed to retrieve user by email: " + email + ", error: " + result.Error.Error())
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByID(id uint) (*usermodel.User, error) {
	var user usermodel.User
	result := ur.db.Table("users").First(&user, id)
	if result.Error != nil {
		ur.logging.Error("Failed to retrieve user by ID: " + fmt.Sprint(id) + ", error: " + result.Error.Error())
		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(user *usermodel.User) error {
	existingUser, err := ur.GetUserByID(user.ID)
	if err != nil {
		ur.logging.Error("Failed to retrieve user by ID: " + err.Error())
		return err
	}

	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.Age = user.Age

	if err := ur.db.Save(existingUser).Error; err != nil {
		ur.logging.Error("Failed to update user: " + err.Error())
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id uint) error {
	result := ur.db.Table("users").Delete(&usermodel.User{}, id)
	if result.Error != nil {
		ur.logging.Error("Failed to delete user with ID: " + fmt.Sprint(id) + ", error: " + result.Error.Error())
		return result.Error
	}
	return nil
}
