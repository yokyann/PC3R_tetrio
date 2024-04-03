package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/yokyann/PC3R_tetrio/server/database"
	. "github.com/yokyann/PC3R_tetrio/server/models"
)

func ListUsers(c *gin.Context) {
	listusers := List_Users()
	if listusers == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "List is empty"})
		return
	}
	//Response to api
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Data fetched successfully", "data": listusers})
}

func CreateUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	id := Create_User(user)
	//end mongo queries
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User Created Successful", "data": id})
}

func FindUser(c *gin.Context) {
	pseudo := c.Param("pseudo")
	user := Find_User(pseudo)
	if user == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	//end mongo queries
	c.IndentedJSON(http.StatusOK, gin.H{"data": user})
}
