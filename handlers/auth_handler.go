package handlers

import (
	"database/sql"
	"net/http"
	"testbbs/internal/auth"
	"testbbs/internal/db"
	"testbbs/internal/util"

	"github.com/golang-jwt/jwt/v5"
)

func RefreshTokenHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		refreshTokenClient, err := r.Cookie("refresh_token")
		if err != nil {
			util.SendErrorResponse(w, http.StatusUnauthorized, "토큰이 없습니다")
			return
		}

		token, err := auth.ValidateToken(refreshTokenClient.Value)
		if err != nil {
			util.SendErrorResponse(w, http.StatusUnauthorized, "Refresh Token이 유효하지 않음")
			return
		}

		// ✅ claims.Claims를 jwt.MapClaims로 변환
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			util.SendErrorResponse(w, http.StatusUnauthorized, "Refresh Token이 유효하지 않음")
			return
		}

		emailVal, ok := claims["email"].(string)
		if !ok {
			util.SendErrorResponse(w, http.StatusUnauthorized, "토큰에서 사용자 정보를 찾을 수 없음")
			return
		}
		email := emailVal

		refreshToken, err := db.SelectRefreshTokensByEmail(database, email)

		// 클라이언트에서 받은 토큰이 실제 DB에 저장되어있는 토큰과 일치하는지 확인
		if err != nil || refreshToken != refreshTokenClient.Value {
			util.SendErrorResponse(w, http.StatusUnauthorized, "Refresh Token이 유효하지 않음")
			return
		}

		// 액세스 토큰 신규 생성
		accessToken, err := auth.GenerateToken(email)
		if err != nil {
			util.SendErrorResponse(w, http.StatusInternalServerError, "토큰 생성 실패")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    accessToken,
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   60 * 15,
		})

		util.SendSuccessResponse(w, http.StatusOK, "액세스 토큰이 갱신되었습니다", nil)
		return

	}
}
