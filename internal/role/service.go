package role

import (
	"context"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type roleService struct {
	collection *mongo.Collection
}

type RoleServiceInterface interface {
	Create(ctx context.Context, roleCreateDto RoleCreateDto) (*Role, error)
}

func NewRoleService(db *mongo.Database) RoleServiceInterface {
	return &roleService{
		collection: db.Collection("roles"),
	}
}

func (s *roleService) Create(ctx context.Context, roleCreateDto RoleCreateDto) (*Role, error) {
	role := &Role{
		Name:   roleCreateDto.Name,
		Description: roleCreateDto.Description,
		PermissionIDs: roleCreateDto.PermissionIDs,
	}

	result, err := s.collection.InsertOne(ctx, role)
	if err != nil {
		logger.Ef("Erreur lors de la création du role : %v", err)
		return nil, err
	}

	role.ID = result.InsertedID.(primitive.ObjectID)

	return role, nil
}