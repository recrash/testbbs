package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"testbbs/internal/auth"
	"testbbs/internal/db"

	"github.com/golang-jwt/jwt/v5"
)

func RefreshTokenHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			RefreshToken string `json:"refresh_token"`
		}

		// 요청 Body에서 RefreshToken 받아오기.
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.RefreshToken == "" {
			http.Error(w, `{"error": "유효하지 않은 검증"}`, http.StatusBadRequest)
			return
		}

		token, err := auth.ValidateToken(req.RefreshToken)
		if err != nil {
			http.Error(w, `{"error": "Refresh Token이 유효하지 않음"}`, http.StatusUnauthorized)
			return
		}

		// ✅ claims.Claims를 jwt.MapClaims로 변환
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, `{"error": "Refresh Token이 유효하지 않음"}`, http.StatusUnauthorized)
			return
		}

		emailVal, ok := claims["email"].(string)
		if !ok {
			http.Error(w, `{"error": "토큰에서 사용자 정보를 찾을 수 없음"}`, http.StatusUnauthorized)
			return
		}
		email := emailVal

		refreshToken, err := db.SelectRefreshTokensByEmail(database, email)

		// 클라이언트에서 받은 토큰이 실제 DB에 저장되어있는 토큰과 일치하는지 확인
		if err != nil || refreshToken != req.RefreshToken {
			http.Error(w, `{"error": "Refresh Token이 유효하지 않음"}`, http.StatusUnauthorized)
			return
		}

		// 액세스 토큰 신규 생성
		accessToken, err := auth.GenerateToken(email)
		if err != nil {
			http.Error(w, `{"error": "토큰 생성 실패"}`, http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

	}
}
