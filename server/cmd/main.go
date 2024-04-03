package main

import (
	"log"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	// router.POST("/books", CreateBook)
	router.GET("/users", ListUsers)
	// router.GET("/books/:name", FindBook)
	router.Run("localhost:8080")
}
