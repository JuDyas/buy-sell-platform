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

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) (string, error)
	Update(ctx context.Context, id primitive.ObjectID, update bson.M) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
}

type categoryRepository struct {
	coll *mongo.Collection
}

func NewCategoryRepository(db *mongo.Database, collection string) CategoryRepository {
	return &categoryRepository{
		coll: db.Collection(collection),
	}
}

func (r *categoryRepository) Create(ctx context.Context, category *models.Category) (string, error) {
	category.ID = primitive.NewObjectID()
	category.CreatedAt = time.Now()
	category.UpdatedAt = category.CreatedAt
	_, err := r.coll.InsertOne(ctx, category)
	if err != nil {
		return "", fmt.Errorf("failed to create category: %w", err)
	}

	return category.ID.Hex(), nil
}

func (r *categoryRepository) Update(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	update["updated_at"] = time.Now()
	_, err := r.coll.UpdateOne(ctx, bson.M{"_id": id, "is_deleted": false}, bson.M{"$set": update})
	if err != nil {
		return fmt.Errorf("failed to update category: %w", err)
	}

	return nil
}

func (r *categoryRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

func (r *categoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Category, error) {
	var category models.Category
	err := r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, fmt.Errorf("failed to find category: %w", err)
	}

	return &category, nil
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	cur, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return []models.Category{}, fmt.Errorf("failed to find categories: %w", err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var category models.Category
		if err := cur.Decode(&category); err == nil {
			categories = append(categories, category)
		}
	}

	if err := cur.Err(); err != nil {
		return []models.Category{}, fmt.Errorf("failed to find categories: %w", err)
	}

	return categories, nil
}
