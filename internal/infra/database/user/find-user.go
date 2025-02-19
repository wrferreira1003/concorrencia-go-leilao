package user_repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/wrferreira1003/concorrencia-go-leilao/config/logger.go"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/entity/user_entity"
	"github.com/wrferreira1003/concorrencia-go-leilao/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepositoryMongo struct {
	Collection *mongo.Collection
}

func NewUserRepositoryMongo(
	collection *mongo.Database,
) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		Collection: collection.Collection("users"),
	}
}

func (r *UserRepositoryMongo) FindUserByID(ctx context.Context, userID string) (*user_entity.User, *internal_error.InternalError) {

	// Define the filter to find the user by ID
	filter := bson.M{"_id": userID}

	var user UserEntityMongo

	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("user not found", err)
			return nil, internal_error.NewNotFoundError("user not found")
		}
		logger.Error("error finding user", err)
		return nil, internal_error.NewInternalServerError("error finding user")
	}

	return &user_entity.User{
		ID:   user.Id,
		Name: user.Name,
	}, nil
}

func (r *UserRepositoryMongo) CreateUser(ctx context.Context, user *user_entity.User) (*user_entity.User, *internal_error.InternalError) {
	user.ID = uuid.New().String()

	_, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		logger.Error("error creating user", err)
		return nil, internal_error.NewInternalServerError("error creating user")
	}

	return user, nil
}
