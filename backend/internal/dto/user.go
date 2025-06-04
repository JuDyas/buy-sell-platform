package dto

import (
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRegister struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}

type UserUpdate struct {
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Location  string `json:"location,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type UserPublic struct {
	ID        primitive.ObjectID `json:"id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	AvatarURL string             `json:"avatar_url"`
	Phone     string             `json:"phone"`
	Location  string             `json:"location"`
	Role      models.UserRole    `json:"role"`
	CreatedAt time.Time          `json:"created_at"`
}
