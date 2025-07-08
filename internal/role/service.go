package role

import (
	"context"
	"eduva-auth/internal/permission"
	"errors"
	"fmt"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type roleService struct {
	collection *mongo.Collection
	permissionService permission.PermissionServiceInterface
}

type RoleServiceInterface interface {
	Create(ctx context.Context, roleCreateDto RoleCreateDto) (*Role, error)
	GetById(ctx context.Context, roleIds []primitive.ObjectID) ([]*Role, error)
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

func (s *roleService) GetById(ctx context.Context, roleIds []primitive.ObjectID) ([]*Role, error) {
	if len(roleIds) == 0 {
		return nil, errors.New("aucun ID fourni")
	}

	filter := bson.M{"_id": bson.M{"$in": roleIds}}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		logger.Ef("Erreur lors de la récupération des roles : %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var roles []*Role

	for cursor.Next(ctx) {
		var role Role
		if err := cursor.Decode(&role); err != nil {
			logger.Ef("Erreur lors du décodage d'un role : %v", err)
			return nil, err
		}
		roles = append(roles, &role)
	}

	if err := cursor.Err(); err != nil {
		logger.Ef("Erreur du curseur lors de la récupération des roles : %v", err)
		return nil, err
	}

	if len(roles) == 0 {
		return nil, errors.New("aucun role trouvé pour les IDs fournis")
	}

	return roles, nil
}