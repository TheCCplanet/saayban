package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sayban/internal/models"
)

// writeJSON writes a JSON response with the given status code and data.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError writes an error response with the given status code and error.
func writeError(w http.ResponseWriter, status int, err error) {
	if err == nil {
		err = fmt.Errorf("unknown error")
	}
	writeJSON(w, status, models.ErrorResponse{Error: err.Error()})
}

// readJSON reads JSON from the request body into the provided value.
func readJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
