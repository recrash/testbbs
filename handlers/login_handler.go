package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testbbs/internal/auth"
	"testbbs/internal/db"
	"testbbs/internal/models"
	"testbbs/internal/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			util.SendErrorResponse(w, http.StatusBadRequest, "잘못된 요청 형식")
			return
		}

		user, err := db.GetUserByEmail(database, req.Email)

		if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
			util.SendErrorResponse(w, http.StatusUnauthorized, "이메일 또는 비밀번호가 잘못되었습니다.")
			return
		}

		// 액세스 토큰 발급
		accessToken, err := auth.GenerateToken(user.Email)
		if err != nil {
			util.SendErrorResponse(w, http.StatusInternalServerError, "토큰 생성 실패")
			return
		}

		refreshToken, err := db.SelectRefreshTokensByEmail(database, user.Email)
		// 행이 없을 경우 신규 리프레시 토큰 생성
		if err != nil {
			var expirationTime time.Time
			refreshToken, expirationTime, err = auth.GenerateRefreshToken(user.Email)
			if err != nil {
				util.SendErrorResponse(w, http.StatusInternalServerError, "리프레시 토큰 생성 실패")
				return
			}
			err = db.InsertRefreshTokens(database, user.Email, refreshToken, expirationTime)
			if err != nil {
				util.SendErrorResponse(w, http.StatusInternalServerError, "리프레시 토큰 DB Insert 실패")
				return
			}
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   60 * 15,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refreshToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   60 * 60 * 24 * 7,
		})

		responseData := map[string]interface{}{
			"user": map[string]string{
				"email":    user.Email,
				"username": user.Username,
			},
		}
		util.SendSuccessResponse(w, http.StatusOK, "로그인 성공!", responseData)

	}
}
