package commentmodel

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Comment represents the Comment table
type Comment struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	PhotoID   uint
	Message   string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate validates the Comment struct
func (c *Comment) Validate() error {
	return validate.Struct(c)
}
