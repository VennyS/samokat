package dto

import "github.com/google/uuid"

type CategoryDTO struct {
	ID       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	ImageURL string         `json:"image_url"`
	Children []*CategoryDTO `json:"children,omitempty"`
}

type CreateCategoryDTO struct {
	Name     string     `json:"name" validate:"required"`
	ImageURL string     `json:"image_url" validate:"required,url"`
	ParentID *uuid.UUID `json:"parent_id,omitempty" validate:"omitempty,uuid"`
}

type UpdateCategoryDTO struct {
	Name     *string    `json:"name,omitempty" validate:"omitempty"`
	ImageURL *string    `json:"image_url,omitempty" validate:"omitempty,url"`
	ParentID *uuid.UUID `json:"parent_id,omitempty" validate:"omitempty,uuid"`
}
