package utils

import (
	"golang_restful_api/models"
	"net/http"
)

// Message ...
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}

// Respond ...
func Respond(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return models.ToJSON(w, data)
}
