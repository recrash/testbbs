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
