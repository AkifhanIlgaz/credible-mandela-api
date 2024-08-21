package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/db"
	"go.mongodb.org/mongo-driver/bson"
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

func (service AdService) DeleteAdById(adId string) error {
	collection := service.db.Collection(db.AdCollection)

	id, err := primitive.ObjectIDFromHex(adId)
	if err != nil {
		return fmt.Errorf("delete ad: %w", err)
	}

	filter := bson.M{
		"_id": id,
	}

	res, err := collection.DeleteOne(service.ctx, filter)
	if err != nil {
		return fmt.Errorf("delete ad: %w", err)
	}

	if res.DeletedCount == 0 {
		return errors.New("ad not found")
	}

	return nil
}

// TODO: Create service functions
