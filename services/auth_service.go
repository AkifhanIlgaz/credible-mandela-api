package services

import (
	"context"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
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

func (service AuthService) CreateUser(user models.User) (models.User, error) {

}
