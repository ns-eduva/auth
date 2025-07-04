package permission

import (
	"context"
	"errors"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type permissionService struct {
	collection *mongo.Collection
}

type PermissionServiceInterface interface {
	Create(ctx context.Context, permissionCreateDto PermissionCreateDto) (*Permission, error)
	GetOneById(ctx context.Context, permissionId primitive.ObjectID) (*Permission, error)
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

func (s *permissionService) GetOneById(ctx context.Context, permissionId primitive.ObjectID) (*Permission, error) {
	var permission Permission

	err := s.collection.FindOne(ctx, bson.M{"_id": permissionId}).Decode(&permission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Ef("Aucune permission trouvée pour l'id : %v", permissionId)
			return nil, errors.New("permission non trouvée")
		}
		logger.Ef("Erreur lors de la récupération de la permission : %v", err)
		return nil, err
	}

	return &permission, nil
}