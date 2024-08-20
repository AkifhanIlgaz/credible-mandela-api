package services

import (
	"context"

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
