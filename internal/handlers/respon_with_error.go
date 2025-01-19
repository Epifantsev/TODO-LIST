package handlers

import (
	"encoding/json"
	"final_task/internal/config"
	"net/http"
)

func ResponWithError(w http.ResponseWriter, status int, answer config.Err) {
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(answer); err != nil {
		http.Error(w, "serialization error", http.StatusInternalServerError)
	}
}
