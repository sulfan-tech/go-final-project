package model

import (
	usermodel "go-final-project/internal/domain/user/model"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type SocialMedia struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Name           string         `validate:"required" json:"name"`
	SocialMediaURL string         `validate:"required" json:"social_media_url"`
	UserID         uint           `json:"user_id"`
	User           usermodel.User `gorm:"foreignKey:UserID" json:"user"` // Embedded User struct
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

func (s *SocialMedia) Validate() error {
	return validate.Struct(s)
}
