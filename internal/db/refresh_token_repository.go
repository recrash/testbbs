package db

import (
	"database/sql"
	"fmt"
	"time"
)

func SelectRefreshTokensByEmail(db *sql.DB, email string) bool {
	var ID int
	query := `SELECT ID FROM refresh_tokens WHERE email = $1 AND expires_at > NOW()`
	err := db.QueryRow(query, email).Scan(ID)

	// row가 없으면 err 발생 -> true 리턴
	if err != nil {
		return true
	}
	// row가 있으면 false 리턴
	return false
}

func InsertRefreshTokens(db *sql.DB, email string, refreshToken string, expiresAt time.Time) error {
	query := `INSERT INTO refresh_tokens(email, token, expires_at) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, email, refreshToken, expiresAt)
	if err != nil {
		return fmt.Errorf("리프레시토큰 Insert 실패: %v", err)
	}

	return nil
}
