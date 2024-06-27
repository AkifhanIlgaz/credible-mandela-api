package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdController struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewAdController(ctx context.Context, db *mongo.Database) (AdController, error) {
	coll := db.Collection("ads")

	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{"address", 1},
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := coll.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return AdController{}, fmt.Errorf("failed to create index: %w", err)
	}

	return AdController{
		ctx:        ctx,
		collection: coll,
	}, nil
}

func (controller AdController) GetAd(ctx *gin.Context) {
	var ad models.Ad

	address := ctx.Query("address")

	err := controller.collection.FindOne(context.Background(), bson.M{"address": address}).Decode(&ad)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &ad)
}

func (controller AdController) GetAllAds(ctx *gin.Context) {
	var ads []models.Ad

	cur, err := controller.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = cur.All(context.TODO(), &ads)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, &ads)
}

func (controller AdController) CreateAd(ctx *gin.Context) {
	address := ctx.Query("address")
	amount, err := strconv.ParseFloat(ctx.Query("amount"), 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ad := models.Ad{
		Address: address,
		Amount:  amount,
	}

	if err := controller.collection.FindOneAndReplace(context.Background(), bson.M{"address": address}, ad).Err(); err == mongo.ErrNoDocuments {
		_, err := controller.collection.InsertOne(controller.ctx, ad)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
}
