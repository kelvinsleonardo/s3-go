package common

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Message  string `json:"message"`
	Datetime string `json:"datetime"`
}

type APIError struct {
	Message    string `json:"message"`
	Datetime   string `json:"datetime"`
	StatusCode int    `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(message string, statusCode int) *APIError {
	return &APIError{
		Message:    message,
		Datetime:   time.Now().Format("2006-01-02 15:04:05"),
		StatusCode: statusCode,
	}
}

func WriteErrorResponse(err error, w http.ResponseWriter) {
	apiErr, ok := err.(*APIError)
	if !ok {
		// Se o erro não for um *APIError personalizado, use um erro genérico
		apiErr = NewAPIError("Erro interno do servidor", http.StatusInternalServerError)
	}

	response := ErrorResponse{
		Message:  apiErr.Message,
		Datetime: apiErr.Datetime,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiErr.StatusCode)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Erro ao serializar resposta:", err)
	}
}
