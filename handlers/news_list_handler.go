package handlers

import (
	"encoding/json"
	"expense-manager-backend/services"
	"net/http"
)

func NewsListHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	languageStr := keys.Get("language")

	newsList, err := services.FetchAndExtractNewsList(languageStr)
	if err != nil {
		http.Error(w, "Failed to fetch news list", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(newsList, "", "  ")
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
