package models

import "time"

type RefreshToken struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Token     string    `json:"Token'`
	ExpiresAt time.Time `json:"expires_at"`
}
