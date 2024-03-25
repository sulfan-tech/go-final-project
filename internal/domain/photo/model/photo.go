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

// Photo represents the Photo table

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `validate:"required"`
	Caption   string
	PhotoURL  string `validate:"required"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time

	// Associations
	User usermodel.User `gorm:"foreignKey:UserID"`
}

// Validate validates the Photo struct
func (p *Photo) Validate() error {
	return validate.Struct(p)
}
