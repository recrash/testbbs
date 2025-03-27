package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testbbs/handlers"
	"testbbs/internal/auth"
	"testbbs/internal/db"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type ResgisterRequest struct {
	Username string `json:username`
	Email    string `json:email`
	Password string `json:password`
}
type LoginRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

func RegisterHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ResgisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "잘못된 요청 형식", http.StatusBadRequest)
			return
		}

		if req.Username == "" || req.Email == "" || req.Password == "" {
			http.Error(w, "모든 필드를 입력하세요.", http.StatusBadRequest)
			return
		}

		err = db.CreateUser(database, req.Username, req.Email, req.Password)
		if err != nil {
			http.Error(w, "회원가입 실패: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "회원가입 성공!")
	}
}

func loginHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

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

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("⚠️  .env 파일을 찾을 수 없습니다. 기본값을 사용합니다.")
	}

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("❌ DB 연결 실패:", err)
	}

	defer database.Close()

	http.HandleFunc("/register", RegisterHandler(database))
	http.HandleFunc("/login", loginHandler(database))
	http.HandleFunc("/profile", auth.AuthMiddleware(handlers.ProfileHandler))
	http.HandleFunc("/refresh", handlers.RefreshTokenHandler(database))

	fmt.Println("🚀 서버 실행 중: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
