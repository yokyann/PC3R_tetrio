// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/yokyann/PC3R_tetrio/server/internal/adapter/handler"
// )

// func main() {
// 	http.HandleFunc("/fetch-username", handler.FetchUsernameHandler)

// 	log.Println("Server is running at http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }


package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yokyann/PC3R_tetrio/server/internal/adapter/database/mongodb"
	"github.com/yokyann/PC3R_tetrio/server/internal/domain"
	"github.com/yokyann/PC3R_tetrio/server/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Get MongoDB URI from environment variable
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	// Connect to MongoDB
	db, err := mongodb.NewDatabase(uri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Close()

	// Initialize user service
	userService := service.NewUserService(db)

	// Example: Query user by pseudo
	pseudo := "yok"
	user, err := userService.GetUserByPseudo(pseudo)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	// Print user details
	log.Printf("User ID: %s\n", user.ID)
	log.Printf("Pseudo: %s\n", user.Pseudo)
	log.Printf("Mail: %s\n", user.Mail)
}
