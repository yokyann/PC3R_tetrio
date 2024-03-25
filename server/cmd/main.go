package main

import (
	"context"
	"log"
	"net/http"

	"github.com/yokyann/PC3R_tetrio/server/internal/adapter/handler"
	"github.com/yokyann/PC3R_tetrio/server/internal/core/database"
)

func main() {
	// Connect to MongoDB
	db, err := database.NewMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer db.Disconnect(context.Background())

	// Set up HTTP request handlers
	http.HandleFunc("/fetch-username", handler.FetchUsernameHandler)

	// Start the server
	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
