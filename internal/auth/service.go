package auth

import (
	"context"
	"errors"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type authService struct {
	collection *mongo.Collection
}

type AuthServiceInterface interface {
	FindByEmail(ctx context.Context, email string) (*Auth, error)
	Create(ctx context.Context, auth *Auth) error
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

func (s *authService) Create(ctx context.Context, auth *Auth) error {
	auth.SetTimeStamps()

	existingAuth, err := s.FindByEmail(ctx, auth.Email)
	if err != nil {
		return err
	}
	if existingAuth != nil {
		logger.Ef("Un utilisateur avec cet email existe déjà : %s", auth.Email)
		return errors.New("impossible de créer votre compte")
	}

	if err := auth.HashPassword(); err != nil {
		return err
	}

	result, err := s.collection.InsertOne(ctx, auth)
	if err != nil {
		logger.Ef("Une erreur est survenue au moment de creer le compte : %v", err)
		return err
	}

	auth.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}