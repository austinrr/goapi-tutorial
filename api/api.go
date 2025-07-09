package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Reponse
type CoinBalanceResponse struct {
	// Success Code, Usually 200
	Code int

	// Account Balance
	Balance int64
}

// Error Response
type Error struct {
	// Error code
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var RequestErrorHandler = func(w http.ResponseWriter, err error) {
	writeError(w, http.StatusBadRequest, err.Error())
}

var InternalErrorHandler = func(w http.ResponseWriter) {
	writeError(w, http.StatusInternalServerError, "Internal server error")
}
