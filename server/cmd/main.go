package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	. "github.com/yokyann/PC3R_tetrio/server/database"
	. "github.com/yokyann/PC3R_tetrio/server/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	defer Disconnect()

	http.HandleFunc("/users", ListUsers)
	http.HandleFunc("/fetch-player", FetchPlayerUsername)

	log.Println("Server listening on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
