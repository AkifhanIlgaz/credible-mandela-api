package models

import "time"

type RefreshToken struct {
	Uid       string
	Token     string
	ExpiresAt time.Time
}
