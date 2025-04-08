package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testbbs/internal/db"
	"testbbs/internal/models"
)

func RegisterHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.ResgisterRequest
		var user *models.User

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, `{"error": "잘못된 요청 형식"}`, http.StatusBadRequest)
			return
		}

		if req.Username == "" || req.Email == "" || req.Password == "" {
			http.Error(w, `{"error": "모든 필드를 입력하세요."}`, http.StatusBadRequest)
			return
		}

		user, _ = db.GetUserByEmail(database, req.Email)
		if user.Email == req.Email {
			http.Error(w, `{"error": "이미 가입된 이메일입니다."}`, http.StatusInternalServerError)
			return
		}

		err = db.CreateUser(database, req.Username, req.Email, req.Password)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "회원가입 실패: %s"}`, err.Error()), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "회원가입 성공!")
	}
}
