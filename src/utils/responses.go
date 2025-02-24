package utils

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse define la estructura para respuestas de error
type ErrorResponse struct {
	Message string `json:"Error"`
}

// RespondWithError envía una respuesta JSON con un mensaje de error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// RespondWithJSON envía una respuesta JSON con datos y código HTTP
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
