package main

import (
	"fmt"
	"log"
	"net/http"
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

	http.HandleFunc("/register", handlers.RegisterHandler(database))
	http.HandleFunc("/login", handlers.LoginHandler(database))
	http.HandleFunc("/profile", auth.AuthMiddleware(handlers.ProfileHandler))
	http.HandleFunc("/refresh", handlers.RefreshTokenHandler(database))

	fmt.Println("🚀 서버 실행 중: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
