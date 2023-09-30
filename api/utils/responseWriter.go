package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWriter(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	// check if data is a simple string
	// if yes, then return it on an object with a key "message"
	// if no, then return the data as it is
	if _, ok := data.(string); ok {
		data = struct {
			Message string `json:"message,omitempty"`
		}{
			Message: data.(string),
		}
	}
	json.NewEncoder(w).Encode(data)
}
