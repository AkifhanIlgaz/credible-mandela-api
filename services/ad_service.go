package services

import (
	"context"

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

// TODO: Create service functions
