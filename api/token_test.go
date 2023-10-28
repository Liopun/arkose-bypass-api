package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	Handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body
	var response SuccessResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	// Validate the response structure and content
	if response.Token == "" {
		t.Error("Handler returned empty token")
	}

	// Check CORS headers
	allowOrigin := rr.Header().Get("Access-Control-Allow-Origin")
	if allowOrigin != "https://leetcode.com" {
		t.Errorf("Wrong Access-Control-Allow-Origin header: got %v want %v", allowOrigin, "https://leetcode.com")
	}

	allowMethods := rr.Header().Get("Access-Control-Allow-Methods")
	if allowMethods != "GET, POST, OPTIONS" {
		t.Errorf("Wrong Access-Control-Allow-Methods header: got %v want %v", allowMethods, "GET, POST, OPTIONS")
	}

	allowHeaders := rr.Header().Get("Access-Control-Allow-Headers")
	if allowHeaders != "Origin, Content-Type" {
		t.Errorf("Wrong Access-Control-Allow-Headers header: got %v want %v", allowHeaders, "Origin, Content-Type")
	}
}
