package handlers

import (
	"encoding/json"
	"net/http"
	"testbbs/internal/auth"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	email, ok := auth.UserFromContext(r.Context())
	if !ok {
		http.Error(w, `{"error": "사용자 정보 없음"}`, http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "프로필 조회 성공!",
		"email":   email,
	})
	return
}
