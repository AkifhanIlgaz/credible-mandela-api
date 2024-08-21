package services

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/credible-mandela-api/config"
	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenService struct {
	config config.Config
}

func NewTokenService(config config.Config) TokenService {
	return TokenService{
		config: config,
	}
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
	claims := models.TokenClaims{
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

	parsedToken, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return jwt.RegisteredClaims{}, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	})

	if err != nil {
		return models.User{}, fmt.Errorf("extract uid from access token: %w", err)
	}

	claims, ok := parsedToken.Claims.(*models.TokenClaims)
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

func (service TokenService) ExtractUserFromRefreshToken(refreshToken string) (models.User, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(service.config.AccessTokenPublicKey)
	if err != nil {
		return models.User{}, fmt.Errorf("extract uid from access token: %w", err)
	}

	parsedToken, err := jwt.ParseWithClaims(refreshToken, &models.TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return jwt.RegisteredClaims{}, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	})

	if err != nil {
		return models.User{}, fmt.Errorf("extract uid from access token: %w", err)
	}

	claims, ok := parsedToken.Claims.(*models.TokenClaims)
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

func (service TokenService) GenerateRefreshToken(uid, address, username string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(service.config.RefreshTokenPrivateKey)
	if err != nil {
		return "", fmt.Errorf("generate refresh token: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("generate refresh token: %w", err)
	}

	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Duration(service.config.RefreshTokenExpiry) * time.Hour))
	claims := models.TokenClaims{
		Username: username,
		Address:  address,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   uid,
			ExpiresAt: expiresAt,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("generate refresh token: %w", err)
	}

	return token, nil
}
