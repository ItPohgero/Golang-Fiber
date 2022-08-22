package models

import (
	"time"
)

type Blog struct {
	ID        uint      `json:"id" validate:"required" gorm:"primaryKey"`
	Slug      string    `json:"slug" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
	UserID    uint      `json:"user_id" validate:"required"`
	User      *User
}
