package dto

import "github.com/google/uuid"

type CategoryDTO struct {
	ID       uuid.UUID      `json:"id"`
	Name     string         `json:"name"`
	ImageURL string         `json:"image_url"`
	Children []*CategoryDTO `json:"children,omitempty"`
}
