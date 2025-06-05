package dto

type AdvertCreate struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Category    string   `json:"category" validate:"required"`
	Price       int      `json:"price" validate:"required"`
	Images      []string `json:"images"`
}

type AdvertUpdate struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Category    string   `json:"category,omitempty"`
	Price       int      `json:"price,omitempty"`
	Images      []string `json:"images,omitempty"`
}
