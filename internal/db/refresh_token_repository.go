package db

import (
	"database/sql"
	"fmt"
	"time"
)

func SelectRefreshTokensByEmail(db *sql.DB, email string) (string, error) {
	var refreshToken string
	query := `SELECT token FROM refresh_tokens WHERE email = $1 AND expires_at > NOW()`
	err := db.QueryRow(query, email).Scan(&refreshToken)

	if err != nil {
		return "", err
	}

	return refreshToken, nil

}

func InsertRefreshTokens(db *sql.DB, email string, refreshToken string, expiresAt time.Time) error {
	queryDelete := `DELETE FROM refresh_tokens WHERE email = $1`
	_, err := db.Exec(queryDelete, email)
	if err != nil {
		return fmt.Errorf("리프레시토큰 Delete 실패: %v", err)
	}

	query := `INSERT INTO refresh_tokens(email, token, expires_at) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, email, refreshToken, expiresAt)
	if err != nil {
		return fmt.Errorf("리프레시토큰 Insert 실패: %v", err)
	}

	return nil
}
