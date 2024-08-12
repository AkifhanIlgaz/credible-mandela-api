package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
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
