package db

import (
	"database/sql"
	"fmt"
	"testbbs/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, username, email, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("사용자 추가 실패: %v", err)
	}

	fmt.Println("✅ 사용자 등록 완료!")
	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	var user models.User

	query := `SELECT * FROM USERS WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
