package dto

type CategoryCreate struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CategoryUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
