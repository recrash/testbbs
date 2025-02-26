package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// JWT를 요청 헤더에서 가져오기
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "토큰이 없습니다"}`, http.StatusUnauthorized)
			return
		}

		// "Bearer 접두사 제거"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, `{"error": "잘못된 토큰 형식"}`, http.StatusUnauthorized)
			return
		}

		// `ValidateToken` 함수로 검증
		token, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, `{"error": "리퀘스트 토큰 검증 에러"}`, http.StatusUnauthorized) // 검증 실패 시, `401 Unauthorized` 응답 반환
			return
		}

		// username에 대한 context 설정
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			r = r.WithContext(WithUserContext(r.Context(), claims["username"].(string)))
		}

		// 검증 성공 시, 다음 핸들러 실행
		next(w, r)
	}
}
