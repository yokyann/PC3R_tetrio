package main

import (
	"log"
	"net/http"

	"github.com/yokyann/PC3R_tetrio/server/internal/adapter/handler/http"
)

func main() {
	http.HandleFunc("/fetch-username", http.FetchUsernameHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
