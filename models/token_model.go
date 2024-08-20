package models

import "time"

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
