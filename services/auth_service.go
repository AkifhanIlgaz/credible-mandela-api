package services

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: All service fields and constructors are duplicate

type AuthService struct {
	ctx context.Context
	db  *mongo.Database
}

func NewAuthService(ctx context.Context, db *mongo.Database) AuthService {
	return AuthService{
		ctx: ctx,
		db:  db,
	}
}

func (service AuthService) Initialize() error {
	collection := service.db.Collection(db.UsersCollection)

	_, err := collection.Indexes().CreateOne(service.ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("initialize auth service: %w", err)
	}

	_, err = collection.Indexes().CreateOne(service.ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "address", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("initialize auth service: %w", err)
	}

	return nil
}

func (service AuthService) CreateUser(user models.User) (primitive.ObjectID, error) {
	collection := service.db.Collection(db.UsersCollection)

	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("create user: %w", err)
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (service AuthService) GetUserByUsername(username string) (models.User, error) {
	collection := service.db.Collection(db.UsersCollection)

	filter := bson.M{
		"username": username,
	}

	var user models.User

	err := collection.FindOne(service.ctx, filter).Decode(&user)
	if err != nil {
		return user, fmt.Errorf("get user by username: %w", err)
	}

	return user, nil
}
