package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testbbs/handlers"
	"testbbs/internal/auth"
	"testbbs/internal/db"

	"github.com/joho/godotenv"
)

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

	port := os.Getenv("SERVER_PORT")

	http.Handle("/register", withCORS(handlers.RegisterHandler(database)))
	http.Handle("/login", withCORS(handlers.LoginHandler(database)))
	http.Handle("/profile", withCORS(auth.AuthMiddleware(handlers.ProfileHandler)))
	http.Handle("/refresh", withCORS(handlers.RefreshTokenHandler(database)))
	http.Handle("/logout", withCORS(handlers.LogOutHandler(database)))

	fmt.Println("🚀 서버 실행 중: http://localhost:8081")
	http.ListenAndServe(port, nil)
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, UPDATE, DELETE, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
