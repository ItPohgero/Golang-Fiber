package models

type Blog struct {
	Id        uint   `json:"id" validate:"required"`
	UserID    uint   `json:"user_id" validate:"required"`
	Slug      string `json:"slug" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Body      string `json:"body" validate:"required"`
	CreatedAt string `json:"created_at" validate:"required"`
	UpdatedAt string `json:"updatedAt" validate:"required"`
}
