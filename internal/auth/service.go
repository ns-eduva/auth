package auth

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type authService struct {
	collection *mongo.Collection
}

type AuthServiceInterface interface {
	// Create(ctx context.Context, auth *Auth) error
}

func NewAuthService(db *mongo.Database) AuthServiceInterface {
	return &authService{
		collection: db.Collection("auth"),
	}
}