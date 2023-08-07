package handlers

import (
	"encoding/json"
	"expense-manager-backend/constants"
	"expense-manager-backend/services"
	"fmt"
	"net/http"
)

func NewsListHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/news/list.php?nt_pk=7", constants.Domain)

	newsList, err := services.FetchAndExtractNewsList(url)
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
