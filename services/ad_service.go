package services

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdService struct {
	ctx context.Context
	db  *mongo.Database
}

func NewAdService(ctx context.Context, db *mongo.Database) AdService {
	return AdService{
		ctx: ctx,
		db:  db,
	}
}

func (service AdService) CreateAd(ad models.Ad) (string, error) {
	collection := service.db.Collection(db.AdCollection)

	res, err := collection.InsertOne(context.Background(), ad)
	if err != nil {
		return "", fmt.Errorf("create ad: %w", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("could not cast inserted ID to ObjectID")
	}

	return id.Hex(), nil
}

// TODO: Create service functions
