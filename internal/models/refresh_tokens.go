package models

import "time"

type RefreshToken struct {
	ID        int       `json:id`
	Username  string    `json:username`
	Token     string    `json:Token`
	ExpiresAt time.Time `json:expires_at`
	Email     string    `json:email`
}
