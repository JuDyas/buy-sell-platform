package repository

import (
	"context"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type AdvertRepository interface {
	Create(ctx context.Context, advert *models.Advert) error
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Advert, error)
	SoftDelete(ctx context.Context, id primitive.ObjectID) error
}

type advertRepository struct {
	coll *mongo.Collection
}

func NewAdvertRepository(db *mongo.Database) AdvertRepository {
	return &advertRepository{
		coll: db.Collection("adverts"),
	}
}

func (r *advertRepository) Create(ctx context.Context, advert *models.Advert) error {
	now := time.Now()
	advert.ID = primitive.NewObjectID()
	advert.CreatedAt = now
	advert.UpdatedAt = now
	advert.IsDeleted = false

	_, err := r.coll.InsertOne(ctx, advert)
	return err
}
