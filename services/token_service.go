package services

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/credible-mandela-api/config"
	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	"github.com/AkifhanIlgaz/credible-mandela-api/utils/db"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenService struct {
	ctx    context.Context
	config config.Config
	db     *mongo.Database
}

func NewTokenService(ctx context.Context, db *mongo.Database, config config.Config) TokenService {
	return TokenService{
		ctx:    ctx,
		db:     db,
		config: config,
	}
}

func (service TokenService) Initialize() error {
	collection := service.db.Collection(db.RefreshTokensCollection)

	_, err := collection.Indexes().CreateOne(service.ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "uid", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("initialize token service: %w", err)
	}

	return nil
}

func (service TokenService) GenerateAccessToken(uid, address, username string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(service.config.AccessTokenPrivateKey)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(service.config.AccessTokenExpiry) * time.Hour))
	claims := models.AccessTokenClaims{
		Username: username,
		Address:  address,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   uid,
			ExpiresAt: expiresAt,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("generate access token: %w", err)
	}

	return token, nil
}

func (service TokenService) ExtractUserFromAccessToken(accessToken string) (models.User, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(service.config.AccessTokenPublicKey)
	if err != nil {
		return models.User{}, fmt.Errorf("extract uid from access token: %w", err)
	}

	parsedToken, err := jwt.ParseWithClaims(accessToken, &models.AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return jwt.RegisteredClaims{}, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	})

	if err != nil {
		return models.User{}, fmt.Errorf("extract uid from access token: %w", err)
	}

	claims, ok := parsedToken.Claims.(*models.AccessTokenClaims)
	if !ok || !parsedToken.Valid {
		return models.User{}, fmt.Errorf("invalid token")
	}

	id, err := primitive.ObjectIDFromHex(claims.Subject)
	if err != nil {
		return models.User{}, fmt.Errorf("extract uid from access token: %w", err)
	}

	user := models.User{
		Username: claims.Username,
		Address:  claims.Address,
		Id:       id,
	}

	return user, nil
}

func (service TokenService) GenerateRefreshToken(uid string) (string, error) {
	refreshToken := models.RefreshToken{
		Uid:       uid,
		Token:     randstr.String(32),
		ExpiresAt: time.Now().Add(time.Duration(service.config.RefreshTokenExpiry) * time.Hour),
	}

	collection := service.db.Collection(db.RefreshTokensCollection)

	_, err := collection.InsertOne(context.Background(), refreshToken)
	if err != nil {
		return "", fmt.Errorf("generate refresh token: %w", err)
	}

	return refreshToken.Token, nil
}

func (service TokenService) RegenerateRefreshToken(token, uid string) (string, error) {
	oldRefreshToken, err := service.getRefreshToken(token)
	if err != nil {
		return "", fmt.Errorf("regenerate refresh token: %w", err)
	}

	if time.Now().After(oldRefreshToken.ExpiresAt) {
		return "", errors.New("refresh token expired")
	}

	newRefreshToken := models.RefreshToken{
		Uid:       uid,
		Token:     randstr.String(32),
		ExpiresAt: time.Now().Add(time.Duration(service.config.RefreshTokenExpiry) * time.Hour),
	}

	collection := service.db.Collection(db.RefreshTokensCollection)

	res, err := collection.ReplaceOne(service.ctx, bson.M{
		"token": token,
	}, newRefreshToken)
	if err != nil {
		return "", fmt.Errorf("regenerate refresh token: %w", err)
	}

	if res.ModifiedCount == 0 {
		return "", fmt.Errorf("token %v not found", token)
	}

	return newRefreshToken.Token, nil
}

func (service TokenService) getRefreshToken(token string) (models.RefreshToken, error) {
	collection := service.db.Collection(db.RefreshTokensCollection)

	var refreshToken models.RefreshToken

	err := collection.FindOne(service.ctx, bson.M{
		"token": token,
	}).Decode(&refreshToken)
	if err != nil {
		return models.RefreshToken{}, fmt.Errorf("get refresh token: %w", err)
	}

	return refreshToken, nil
}

func (service TokenService) DeleteRefreshToken(token string) error {
	collection := service.db.Collection(db.RefreshTokensCollection)

	res, err := collection.DeleteOne(service.ctx, bson.M{
		"token": token,
	})
	if err != nil {
		return fmt.Errorf("delete refresh token: %w", err)
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("refresh token %v does not exist in database", token)
	}

	return nil
}
