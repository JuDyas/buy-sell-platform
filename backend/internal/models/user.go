package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRole int

const (
	RoleUser UserRole = iota
	RoleAdmin
	RoleModerator
)

func (r UserRole) String() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleModerator:
		return "moderator"
	default:
		return "user"
	}
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	CreatedAt    time.Time          `bson:"created_at"`
	AvatarURL    string             `bson:"avatar_url,omitempty"`
	Phone        string             `bson:"phone,omitempty"`
	Location     string             `bson:"location,omitempty"`
	Role         UserRole           `bson:"role"`
}
