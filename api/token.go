package handler

import (
	"encoding/json"
	"net/http"

	"github.com/flyingpot/funcaptcha"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Token string `json:"token"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Attempt to get the token
	token, err := funcaptcha.GetOpenAITokenV2()
	if err != nil {
		// Handle the error and send an error response
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := ErrorResponse{Error: err.Error()}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Send a success response with the token
	w.WriteHeader(http.StatusOK)
	successResponse := SuccessResponse{Token: token}
	json.NewEncoder(w).Encode(successResponse)
}
