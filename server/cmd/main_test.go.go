package main

import (
	"context"
	"log"
	"testing"
)

func TestMongoDBConnection(t *testing.T) {
	// Définir l'URI de connexion à MongoDB
	uri := "mongodb+srv://yokyann:U2x6Tp7ZRzMvrO7b@tetrio.hdafk7v.mongodb.net/?retryWrites=true&w=majority&appName=tetrio"

	// Créer un contexte
	ctx := context.TODO()

	// Établir une connexion à MongoDB
	db, err := NewMongoDB(ctx, uri)
	if err != nil {
		t.Errorf("Erreur lors de la connexion à MongoDB : %v", err)
		return
	}

	// Déconnecter la base de données MongoDB
	defer func() {
		if err := db.client.Disconnect(ctx); err != nil {
			log.Fatalf("Erreur lors de la déconnexion de MongoDB : %v", err)
		}
	}()

	// Vérifier si la connexion est établie avec succès
	if err := db.client.Ping(ctx, nil); err != nil {
		t.Errorf("Erreur lors de la vérification de la connexion à MongoDB : %v", err)
		return
	}

	// La connexion à MongoDB a réussi
	t.Logf("Connexion à MongoDB réussie !")
}
