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

type CommunityNoteService struct {
	ctx context.Context
	db  *mongo.Database
}

func NewCommunityNoteService(ctx context.Context, db *mongo.Database) CommunityNoteService {
	return CommunityNoteService{
		ctx: ctx,
		db:  db,
	}
}

func (service CommunityNoteService) Create(communityNote models.CommunityNote) (string, error) {
	collection := service.db.Collection(db.CommunityNotesCollection)

	res, err := collection.InsertOne(context.Background(), communityNote)
	if err != nil {
		return "", fmt.Errorf("create community note: %w", err)
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("could not cast inserted ID to ObjectID")
	}

	return id.Hex(), nil
}

func (service CommunityNoteService) DeleteById(communityNoteId string) error {
	collection := service.db.Collection(db.CommunityNotesCollection)

	id, err := primitive.ObjectIDFromHex(communityNoteId)
	if err != nil {
		return fmt.Errorf("delete community note: %w", err)
	}

	filter := bson.M{
		"_id": id,
	}

	res, err := collection.DeleteOne(service.ctx, filter)
	if err != nil {
		return fmt.Errorf("delete community note: %w", err)
	}

	if res.DeletedCount == 0 {
		return errors.New("community note not found")
	}

	return nil
}

func (service CommunityNoteService) GetById(communityNoteId string) (models.CommunityNote, error) {
	collection := service.db.Collection(db.CommunityNotesCollection)

	id, err := primitive.ObjectIDFromHex(communityNoteId)
	if err != nil {
		return models.CommunityNote{}, fmt.Errorf("get community note by id: %w", err)
	}

	filter := bson.M{
		"_id": id,
	}

	var communityNote models.CommunityNote

	err = collection.FindOne(service.ctx, filter).Decode(&communityNote)
	if err != nil {
		return models.CommunityNote{}, fmt.Errorf("get community note by id: %w", err)
	}

	// Get isLiked, likeCount, CRED

	return communityNote, nil
}

func (service CommunityNoteService) GetAll() ([]models.CommunityNote, error) {
	collection := service.db.Collection(db.CommunityNotesCollection)

	cursor, err := collection.Find(service.ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("get all community notes: %w", err)
	}

	communityNotes := []models.CommunityNote{}

	err = cursor.All(service.ctx, &communityNotes)
	if err != nil {
		return nil, fmt.Errorf("get all community notes: %w", err)
	}

	return communityNotes, nil
}

func (service CommunityNoteService) GetByUsername(username string) ([]models.CommunityNote, error) {
	collection := service.db.Collection(db.CommunityNotesCollection)

	filter := bson.M{
		"publisher": username,
	}

	cursor, err := collection.Find(service.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("get all community notes of user: %w", err)
	}

	communityNotes := []models.CommunityNote{}

	err = cursor.All(service.ctx, &communityNotes)
	if err != nil {
		return nil, fmt.Errorf("get all community notes of user: %w", err)
	}

	return communityNotes, nil
}
