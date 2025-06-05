package dto

type CategoryCreate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
