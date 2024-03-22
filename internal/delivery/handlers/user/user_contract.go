package user

import (
	usermodel "go-final-project/internal/domain/user/model"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (u *LoginRequest) Validate() error {
	return validate.Struct(u)
}

type RegistrationRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `validate:"required,min=6" json:"password,omitempty"`
	Age      int    `json:"age" validate:"required,min=9"`
}

func (u *RegistrationRequest) Validate() error {
	return validate.Struct(u)
}

type LoginResponse struct {
	Token string         `json:"token"`
	User  usermodel.User `json:"user"`
}

type RegistrationResponse struct {
	Message string `json:"message"`
}
