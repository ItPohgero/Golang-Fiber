package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email,min=6,max=32"`
	Password  []byte    `json:"-" validate:"required"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime:true"`
	Blogs     []Blog    `json:"blogs" gorm:"foreignKey:UserID"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
