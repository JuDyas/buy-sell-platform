package repository

import (
	"context"
	"errors"
	"github.com/JuDyas/buy-sell-platform/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) UserRepository {
	return &userRepository{
		collection: db.Collection(collectionName),
	}
}

func (ur *userRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	_, err := ur.collection.InsertOne(ctx, user)

	return err
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := ur.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &user, err
}

func (ur *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := ur.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &user, err
}

func (ur *userRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := ur.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}

	return &user, err
}
