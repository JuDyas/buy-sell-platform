package repository

import (
	"context"
	"fmt"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type AdvertRepository interface {
	Create(ctx context.Context, advert *models.Advert) (*models.Advert, error)
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) error
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.Advert, error)
	SoftDelete(ctx context.Context, id primitive.ObjectID) error
	GetAll(ctx context.Context) ([]models.Advert, error)
	GetByCategory(ctx context.Context, categoryId primitive.ObjectID) ([]models.Advert, error)
	GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Advert, error)
}

type advertRepository struct {
	coll *mongo.Collection
}

func NewAdvertRepository(db *mongo.Database, collection string) AdvertRepository {
	return &advertRepository{
		coll: db.Collection(collection),
	}
}

func (r *advertRepository) Create(ctx context.Context, advert *models.Advert) (*models.Advert, error) {
	now := time.Now()
	advert.ID = primitive.NewObjectID()
	advert.CreatedAt = now
	advert.UpdatedAt = now
	advert.IsDeleted = false

	_, err := r.coll.InsertOne(ctx, advert)
	return advert, err
}

func (r *advertRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	update["updated_at"] = time.Now()
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": id, "is_deleted": false}, bson.M{"$set": update})

	return err
}

func (r *advertRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.Advert, error) {
	var advert models.Advert
	err := r.coll.FindOne(ctx, bson.M{"_id": id, "is_deleted": false}).Decode(&advert)
	if err != nil {
		return nil, err
	}

	return &advert, nil
}

func (r *advertRepository) SoftDelete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.coll.UpdateOne(
		ctx,
		bson.M{"_id": id, "is_deleted": false},
		bson.M{"$set": bson.M{"is_deleted": true, "updated_at": time.Now()}},
	)
	return err
}

func (r *advertRepository) GetAll(ctx context.Context) ([]models.Advert, error) {
	cur, err := r.coll.Find(ctx, bson.M{"is_deleted": false})
	if err != nil {
		return nil, fmt.Errorf("failed to find adverts: %w", err)
	}
	defer cur.Close(ctx)

	var adverts []models.Advert
	for cur.Next(ctx) {
		var advert models.Advert
		if err := cur.Decode(&advert); err == nil {
			adverts = append(adverts, advert)
		}
	}

	return adverts, nil
}

func (r *advertRepository) GetByCategory(ctx context.Context, categoryId primitive.ObjectID) ([]models.Advert, error) {
	cur, err := r.coll.Find(ctx, bson.M{
		"category_id": categoryId,
		"is_deleted":  false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find adverts: %w", err)
	}
	defer cur.Close(ctx)

	var adverts []models.Advert
	for cur.Next(ctx) {
		var advert models.Advert
		if err := cur.Decode(&advert); err == nil {
			adverts = append(adverts, advert)
		}
	}

	return adverts, nil
}

func (r *advertRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Advert, error) {
	cur, err := r.coll.Find(ctx, bson.M{
		"author_id":  userID,
		"is_deleted": false,
	})
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			return
		}
	}(cur, ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to find adverts: %w", err)
	}

	var adverts []models.Advert
	for cur.Next(ctx) {
		var advert models.Advert
		if err := cur.Decode(&advert); err == nil {
			adverts = append(adverts, advert)
		}
	}

	return adverts, nil
}
