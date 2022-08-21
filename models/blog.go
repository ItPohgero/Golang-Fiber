package models

import "time"

type Blog struct {
	Id        uint      `json:"id" validate:"required"`
	UserID    uint      `json:"user_id" validate:"required"`
	Slug      string    `json:"slug" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}
