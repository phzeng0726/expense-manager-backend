package handlers

import (
	"encoding/json"
	"expense-manager-backend/services"
	"net/http"
)

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	idStr := keys.Get("id")

	news, err := services.FetchAndExtractNews(idStr)
	if err != nil {
		http.Error(w, "Failed to fetch news details", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(news, "", "  ")
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
