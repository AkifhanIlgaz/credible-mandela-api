package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ad struct {
	Id         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Advertiser string             `json:"advertiser" bson:"advertiser"` // address
	Amount     float64            `json:"amount" bson:"amount"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
}

type PublishAdForm struct {
	Amount float64 `json:"amount" binding:"required"`
}

type PublishAdResponse struct {
	AdId string `json:"adId"`
}

type UpdateAdForm struct {
	NewAmount float64 `json:"newAmount" binding:"required"`
}
