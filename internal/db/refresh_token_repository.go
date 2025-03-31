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
	// 기존 refresh Token 삭제
	err := DeleteRefreshTokensAll(db, email)

	query := `INSERT INTO refresh_tokens(email, token, expires_at) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, email, refreshToken, expiresAt)
	if err != nil {
		return fmt.Errorf("리프레시토큰 Insert 실패: %v", err)
	}

	return nil
}

func DeleteRefreshTokensAll(db *sql.DB, email string) error {
	queryDelete := `DELETE FROM refresh_tokens WHERE email = $1`
	_, err := db.Exec(queryDelete, email)
	if err != nil {
		return fmt.Errorf("failed DeleteRefreshTokensAll: %v", err)
	}

	return nil
}

func DeleteRefreshTokensByEmail(db *sql.DB, email string, refreshToken string) error {
	queryDelete := `DELETE FROM refresh_tokens WHERE email = $1 AND token = $2`
	_, err := db.Exec(queryDelete, email, refreshToken)
	if err != nil {
		return fmt.Errorf("failed DeleteRefreshTokensByEmail: %v", err)
	}

	return nil
}
