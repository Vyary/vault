package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"vault/internal/models"
)

func Error(w http.ResponseWriter, statusCode int, message string, err error, validationErrors map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errResp := models.ErrorResponse{
		Code:    statusCode,
		Message: message,
	}

	if len(validationErrors) > 0 {
		errResp.Errors = validationErrors
	}

	if err != nil {
		slog.Warn(message, "error", err)
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
