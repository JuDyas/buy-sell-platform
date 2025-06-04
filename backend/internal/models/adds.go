package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Advert struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       int                `bson:"price"`
	Images      []string           `bson:"images"`
	AuthorID    primitive.ObjectID `bson:"author_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	IsDeleted   bool               `bson:"is_deleted"`
}
