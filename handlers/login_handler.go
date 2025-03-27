package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testbbs/internal/auth"
	"testbbs/internal/db"
	"testbbs/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, `{"error": "잘못된 요청 형식"}`, http.StatusBadRequest)
			return
		}

		user, err := db.GetUserByEmail(database, req.Email)

		if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
			http.Error(w, `{"error": "이메일 또는 비밀번호가 잘못되었습니다."}`, http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(user.Email)
		if err != nil {
			http.Error(w, `{"error": "토큰 생성 실패"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Authorization", "Bearer "+token) // 헤더에 JWT 포함
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("로그인 성공! 환영합니다. %s!", user.Username),
			"token":   token,
		})
	}
}
