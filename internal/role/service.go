package role

import (
	"context"
	"eduva-auth/internal/permission"
	"fmt"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type roleService struct {
	collection *mongo.Collection
	permissionService permission.PermissionServiceInterface
}

type RoleServiceInterface interface {
	Create(ctx context.Context, roleCreateDto RoleCreateDto) (*Role, error)
}

func NewRoleService(db *mongo.Database, permissionService permission.PermissionServiceInterface) RoleServiceInterface {
	return &roleService{
		collection:        db.Collection("roles"),
		permissionService: permissionService,
	}
}

func (s *roleService) Create(ctx context.Context, roleCreateDto RoleCreateDto) (*Role, error) {
	validIDs, err := s.permissionService.FilterExisting(ctx, roleCreateDto.PermissionIDs)
	if err != nil {
		return nil, err
	}

	if len(validIDs) == 0 {
		return nil, fmt.Errorf("aucune permission valide trouvée")
	}

	role := &Role{
		Name:   roleCreateDto.Name,
		Description: roleCreateDto.Description,
		PermissionIDs: validIDs,
	}

	result, err := s.collection.InsertOne(ctx, role)
	if err != nil {
		logger.Ef("Erreur lors de la création du role : %v", err)
		return nil, err
	}

	role.ID = result.InsertedID.(primitive.ObjectID)

	return role, nil
}