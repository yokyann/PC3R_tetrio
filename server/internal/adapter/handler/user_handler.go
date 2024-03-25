package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/yokyann/PC3R_tetrio/server/internal/core/service"
)

func FetchUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le nom d'utilisateur de la requête
	username := r.URL.Query().Get("username")

	// Appeler la fonction de service pour récupérer le nom d'utilisateur
	fetchedUsername, err := user_service.FetchUsername(username)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch username: %v", err), http.StatusInternalServerError)
		return
	}

	// Renvoyer le nom d'utilisateur dans la réponse
	jsonResponse, err := json.Marshal(map[string]string{"username": fetchedUsername})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal JSON response: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

