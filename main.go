package main

import (
	"expense-manager-backend/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes.SetupRoutes()

	port := "8080"
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
