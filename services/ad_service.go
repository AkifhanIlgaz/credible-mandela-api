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

func (service AdService) Create(ad models.Ad) (string, error) {
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

func (service AdService) GetAds() ([]models.Ad, error) {
	collection := service.db.Collection(db.AdCollection)

	cursor, err := collection.Find(service.ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("get ads: %w", err)
	}

	ads := []models.Ad{}

	err = cursor.All(service.ctx, &ads)
	if err != nil {
		return nil, fmt.Errorf("get ads: %w", err)
	}

	return ads, nil
}

func (service AdService) GetAdsByAddress(address string) ([]models.Ad, error) {
	collection := service.db.Collection(db.AdCollection)

	filter := bson.M{
		"advertiser": address,
	}

	cursor, err := collection.Find(service.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get ads by address: %w", err)
	}

	ads := []models.Ad{}

	err = cursor.All(service.ctx, &ads)
	if err != nil {
		return nil, fmt.Errorf("get ads by address: %w", err)
	}

	return ads, nil
}

func (service AdService) GetById(adId string) (models.Ad, error) {
	collection := service.db.Collection(db.AdCollection)

	id, err := primitive.ObjectIDFromHex(adId)
	if err != nil {
		return models.Ad{}, fmt.Errorf("get ad by id: %w", err)
	}

	filter := bson.M{
		"_id": id,
	}

	var ad models.Ad

	err = collection.FindOne(service.ctx, filter).Decode(&ad)
	if err != nil {
		return models.Ad{}, fmt.Errorf("get ad by id: %w", err)
	}

	return ad, nil
}

func (service AdService) DeleteById(adId string) error {
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

func (service AdService) UpdateById(adId string, newAmount float64) error {
	collection := service.db.Collection(db.AdCollection)

	id, err := primitive.ObjectIDFromHex(adId)
	if err != nil {
		return fmt.Errorf("delete ad: %w", err)
	}

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"amount": newAmount,
		},
	}

	result, err := collection.UpdateOne(service.ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("ad not found")
	}

	return nil
}
