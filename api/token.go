package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/flyingpot/funcaptcha"
)

type Response struct {
	Token string `json:"message,omitempty"`
	Error string `json:"error,omitempty"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := make(map[string]string)

	token, err := funcaptcha.GetOpenAITokenV2()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response["error"] = err.Error()
	} else {
		w.WriteHeader(http.StatusOK)
		response["token"] = token
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error happened in JSON marshal. Err:", err)
	} else {
		w.Write(jsonResponse)
	}
}
