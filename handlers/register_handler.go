package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"testbbs/internal/db"
	"testbbs/internal/models"
	"testbbs/internal/util"
)

func RegisterHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			util.SendErrorResponse(w, http.StatusBadRequest, "잘못된 요청 형식")
			return
		}

		if req.Username == "" || req.Email == "" || req.Password == "" {
			util.SendErrorResponse(w, http.StatusBadRequest, "모든 필드를 입력하세요.")
			return
		}

		_, err = db.GetUserByEmail(database, req.Email)
		if err == nil {
			util.SendErrorResponse(w, http.StatusBadRequest, "이미 가입된 이메일주소입니다.")
			return
		} else if errors.Is(err, sql.ErrNoRows) {
			err = db.CreateUser(database, req.Username, req.Email, req.Password)
			if err != nil {
				util.SendErrorResponse(w, http.StatusInternalServerError, "회원가입에 실패했습니다. 잠시 후 다시 시도하세요.")
				log.Println("에러 발생 상세: ", err)
				return
			}
			util.SendSuccessResponse(w, http.StatusCreated, "회원가입 성공", nil)
			return
		} else {
			util.SendErrorResponse(w, http.StatusInternalServerError, "사용자 조회 도중 에러가 발생하였습니다. 잠시 후 다시 시도하세요.")
			log.Println("에러 발생 상세: ", err)
			return
		}
	}
}
