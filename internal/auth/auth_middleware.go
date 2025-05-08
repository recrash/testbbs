package auth

import (
	"net/http"
	"testbbs/internal/util"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// JWT를 요청 헤더에서 가져오기
		accessToken, err := r.Cookie("access_token")
		if err != nil {
			util.SendErrorResponse(w, http.StatusUnauthorized, "토큰이 없습니다")
			return
		}

		// `ValidateToken` 함수로 검증
		token, err := ValidateToken(accessToken.Value)
		if err != nil {
			util.SendErrorResponse(w, http.StatusUnauthorized, "리퀘스트 토큰 검증 에러")
			return
		}

		// username에 대한 context 설정
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r = r.WithContext(WithUserContext(r.Context(), claims["email"].(string)))
		}

		// 검증 성공 시, 다음 핸들러 실행
		next(w, r)
	}
}
