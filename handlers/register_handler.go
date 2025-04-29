package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"testbbs/internal/db"
	"testbbs/internal/models"
	"testbbs/internal/util"
)

func RegisterHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest
		var user *models.User

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			util.SendErrorResponse(w, http.StatusBadRequest, "잘못된 요청 형식")
			return
		}

		if req.Username == "" || req.Email == "" || req.Password == "" {
			util.SendErrorResponse(w, http.StatusBadRequest, "모든 필드를 입력하세요.")
			return
		}

		user, _ = db.GetUserByEmail(database, req.Email)
		if user.Email == req.Email {
			util.SendErrorResponse(w, http.StatusInternalServerError, "이미 가입된 이메일주소입니다.")
			return
		}

		err = db.CreateUser(database, req.Username, req.Email, req.Password)
		if err != nil {
			util.SendErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf(`{"error": "회원가입 실패: %s"}`, err.Error()))
			return
		}

		util.SendSuccessResponse(w, http.StatusCreated, "회원가입 성공", nil)
		fmt.Fprintln(w, "회원가입 성공!")
	}
}
