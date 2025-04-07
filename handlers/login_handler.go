package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testbbs/internal/auth"
	"testbbs/internal/db"
	"testbbs/internal/models"
	"time"

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

		// 액세스 토큰 발급
		accessToken, err := auth.GenerateToken(user.Email)
		if err != nil {
			http.Error(w, `{"error": "토큰 생성 실패"}`, http.StatusInternalServerError)
			return
		}

		refreshToken, err := db.SelectRefreshTokensByEmail(database, user.Email)
		// 행이 없을 경우 신규 리프레시 토큰 생성
		if err != nil {
			var expirationTime time.Time
			refreshToken, expirationTime, err = auth.GenerateRefreshToken(user.Email)
			if err != nil {
				http.Error(w, `{"error": "리프레시 토큰 생성 실패"}`, http.StatusInternalServerError)
				return
			}
			err = db.InsertRefreshTokens(database, user.Email, refreshToken, expirationTime)
			if err != nil {
				http.Error(w, `{"error": "리프레시 토큰 DB Insert 실패"}`, http.StatusInternalServerError)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Authorization", "Bearer "+accessToken) // 헤더에 JWT 포함
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

	}
}
