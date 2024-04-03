package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	. "github.com/yokyann/PC3R_tetrio/server/database"
	. "github.com/yokyann/PC3R_tetrio/server/routes"
)

func main() {
	// Charger les variables d'environnement
	err := LoadEnv()
	if err != nil {
		log.Fatalln(err)
	}

	defer Disconnect()

	http.HandleFunc("/users", ListUsers)
	// Ajoutez d'autres routes ici au besoin

	log.Println("Server listening on localhost:8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// LoadEnv charge les variables d'environnement à partir des variables d'environnement système
func LoadEnv() error {
	mongodbURI := os.Getenv("MONGODB_URI")
	log.Println("MONGODB_URI:", mongodbURI) // Ajoutez cette ligne pour vérifier la valeur de MONGODB_URI
	if mongodbURI == "" {
		return errors.New("MONGODB_URI environment variable not set")
	}
	// Chargez d'autres variables d'environnement au besoin

	return nil
}
