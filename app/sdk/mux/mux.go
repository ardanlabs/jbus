package mux

import (
	"encoding/json"
	"net/http"
)

func WebAPI() http.Handler {
	mux := http.NewServeMux()

	h := func(w http.ResponseWriter, r *http.Request) {
		status := struct {
			Status string
		}{
			Status: "OK",
		}

		json.NewEncoder(w).Encode(status)
	}

	mux.HandleFunc("GET /test", h)

	return mux
}
