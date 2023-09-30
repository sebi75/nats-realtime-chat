package ping

import (
	"api/utils/logger"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received ping request...")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "pong"}`))
}
