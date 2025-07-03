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
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
}

func NewAuthService(db *mongo.Database) AuthServiceInterface {
	return &authService{
		collection: db.Collection("users"),
	}
}

func (s *authService) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := s.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Wf("Erreur mongo no document: %v", err)
			return nil, nil
		}
		logger.Ef("Erreur à la recuperation de l'utilisateur: %v", err)
		return nil, err
	}
	return &user, nil
}

func (s *authService) Create(ctx context.Context, user *User) error {
	user.SetTimeStamps()

	existingUser, err := s.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		logger.Ef("Un utilisateur avec cet email existe déjà : %s", user.Email)
		return errors.New("impossible de créer votre compte")
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		logger.Ef("Une erreur est survenue au moment de creer le compte : %v", err)
		return err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}