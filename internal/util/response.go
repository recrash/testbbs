package util

import (
	"encoding/json"
	"net/http"
	"testbbs/internal/models"
)

func SendSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	response := models.ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
	response := models.ApiResponse{
		Success: false,
		Error:   errorMsg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
