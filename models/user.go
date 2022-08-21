package models

type User struct {
	Id       uint   `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password []byte `json:"-" validate:"required"`
}
