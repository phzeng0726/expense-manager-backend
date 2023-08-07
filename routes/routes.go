package routes

import (
	"expense-manager-backend/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/news", handlers.NewsHandler)
	http.HandleFunc("/news-list", handlers.NewsListHandler)
}
