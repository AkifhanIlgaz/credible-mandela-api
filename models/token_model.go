package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type RefreshToken struct {
	Uid       string
	Token     string
	ExpiresAt time.Time
}

type RefreshTokenForm struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AccessTokenClaims struct {
	Username string `json:"username"`
	Address  string `json:"address"`
	jwt.RegisteredClaims
}
