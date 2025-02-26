package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	// 서버 키 가져오기
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	// 클레임 생성(유저이름 / 만료시간)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// 서버 키와 클레임을 이용하여 JWT 생성(서명은 없음) -> 헤더(Header).페이로드(Payload).(서명 없음)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 서버키를 이용하여 서명까지 한 토큰을 리턴
	return token.SignedString(secretKey)
}

// 1️⃣ 클라이언트가 보낸 JWT를 받음
// 2️⃣ 서명이 올바른지 확인 (비밀 키로 검증)
// 3️⃣ JWT가 만료되지 않았는지 확인 (exp 체크)
// 4️⃣ 검증이 완료되면 사용자 정보를 반환
func ValidateToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return nil, errors.New("토큰이 만료되었습니다")
		}
	}

	return token, nil

}
