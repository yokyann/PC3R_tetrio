package connect

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB représente la structure de connexion à la base de données MongoDB.
type MongoDB struct {
    client *mongo.Client
    // Vous pouvez ajouter d'autres champs ici selon vos besoins.
}

// NewMongoDB crée une nouvelle instance de MongoDB.
func NewMongoDB(ctx context.Context, uri string) (*MongoDB, error) {
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }

    // Vérifier si la connexion à la base de données est réussie.
    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    return &MongoDB{
        client: client,
    }, nil
}

// Close ferme la connexion à la base de données MongoDB.
func (db *MongoDB) Close(ctx context.Context) error {
    return db.client.Disconnect(ctx)
}
