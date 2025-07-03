package auth

import (
	"context"
	"errors"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	collection *mongo.Collection
}

type UserServiceInterface interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, userCreateDto UserCreateDto) (*User, error)
}

func NewUserService(db *mongo.Database) UserServiceInterface {
	return &userService{
		collection: db.Collection("users"),
	}
}

func (s *userService) FindByEmail(ctx context.Context, email string) (*User, error) {
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

func (s *userService) Create(ctx context.Context, userCreateDto UserCreateDto) (*User, error) {
	user := &User{
		Email:   userCreateDto.Email,
		Password: userCreateDto.Password,
		RoleIDs: userCreateDto.RoleIDs,
	}

	user.SetTimeStamps()

	existingUser, err := s.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		logger.Ef("Un utilisateur avec cet email existe déjà : %s", user.Email)
		return nil, errors.New("impossible de créer votre compte")
	}

	if err := user.HashPassword(); err != nil {
		return nil, err
	}

	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		logger.Ef("Erreur lors de la création de l'utilisateur : %v", err)
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}
