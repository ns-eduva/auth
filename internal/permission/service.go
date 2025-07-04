package permission

import (
	"context"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type permissionService struct {
	collection *mongo.Collection
}

type PermissionServiceInterface interface {
	Create(ctx context.Context, permissionCreateDto PermissionCreateDto) (*Permission, error)
}

func NewPermissionService(db *mongo.Database) PermissionServiceInterface {
	return &permissionService{
		collection: db.Collection("permissions"),
	}
}

func (s *permissionService) Create(ctx context.Context, permissionCreateDto PermissionCreateDto) (*Permission, error) {
	permission := &Permission{
		Name:   permissionCreateDto.Name,
		Description: permissionCreateDto.Description,
	}

	result, err := s.collection.InsertOne(ctx, permission)
	if err != nil {
		logger.Ef("Erreur lors de la création de la permission : %v", err)
		return nil, err
	}

	permission.ID = result.InsertedID.(primitive.ObjectID)

	return permission, nil
}
