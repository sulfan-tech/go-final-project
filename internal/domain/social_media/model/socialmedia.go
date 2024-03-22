package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `validate:"required"`
	SocialMediaURL string `validate:"required"`
	UserID         uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (s *SocialMedia) Validate() error {
	return validate.Struct(s)
}
