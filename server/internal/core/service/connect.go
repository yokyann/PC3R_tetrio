package service

import (
    "context"
    "github.com/yokyann/PC3R_tetrio/server/internal/adapter/repository"
)

// Exemple d'utilisation de la connexion MongoDB dans un service.
type UserService struct {
    db *repository.MongoDB
}

// NewUserService crée une nouvelle instance de UserService.
func NewUserService(db *repository.MongoDB) *UserService {
    return &UserService{
        db: db,
    }
}

// Vous pouvez ajouter d'autres méthodes à UserService pour interagir avec MongoDB.
