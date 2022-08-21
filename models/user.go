package models

import "time"

type User struct {
	ID        uint      `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`
	Password  []byte    `json:"-" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}
