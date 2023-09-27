package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWriter(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
