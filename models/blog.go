package models

import (
	"time"
)

type Blog struct {
	ID        uint      `json:"id" validate:"required"`
	UserID    uint      `json:"user_id" validate:"required"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Slug      string    `json:"slug" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Body      string    `json:"body" validate:"required"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
