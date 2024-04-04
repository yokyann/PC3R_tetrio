// routes/routes.go
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/yokyann/PC3R_tetrio/server/handler"
	"github.com/yokyann/PC3R_tetrio/server/models"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := handler.List_Users()
	if users == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "List is empty"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "Data fetched successfully", "data": users})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad request"})
		return
	}
	id := handler.Create_User(user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "User Created Successful", "data": id})
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	pseudo := r.URL.Query().Get("pseudo")
	user := handler.Find_User(pseudo)
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "User not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"data": user})
}
