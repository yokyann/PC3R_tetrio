package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yokyann/PC3R_tetrio/server/routes"
)

func FetchUsernameHandler(w http.ResponseWriter, r *http.Request) {
	// Extraire le paramètre "username" de la requête HTTP
	username := r.URL.Query().Get("username")

	// Appeler la fonction FetchPlayerUsername avec le paramètre "username"
	username, err := routes.FetchPlayerUsername(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Écrire la réponse JSON avec le username récupéré
	response := map[string]string{"username": username}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
