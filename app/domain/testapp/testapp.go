package testapp

import (
	"encoding/json"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	// Unmarshal Input Data
	// Validate Data
	// Business Layer
	// Error: return error
	// Success: Form data reponse

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
}
