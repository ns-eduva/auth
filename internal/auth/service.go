package auth

import (
	"context"
	"errors"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authService struct {
	collection *mongo.Collection
}

type AuthServiceInterface interface {
	FindByEmail(ctx context.Context, email string) (*Auth, error)
}

func NewAuthService(db *mongo.Database) AuthServiceInterface {
	return &authService{
		collection: db.Collection("auth"),
	}
}

func (s *authService) FindByEmail(ctx context.Context, email string) (*Auth, error) {
	var auth Auth
	err := s.collection.FindOne(ctx, bson.M{"email": email}).Decode(&auth)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Wf("Erreur mongo no document: %v", err)
			return nil, nil
		}
		logger.Ef("Erreur à la recuperation du auth: %v", err)
		return nil, err
	}
	return &auth, nil
}