package role

import (
	"eduva-auth/internal/permission"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name          string               `bson:"name" json:"name"`
	Description   string               `bson:"description" json:"description"`
	PermissionIDs []primitive.ObjectID `bson:"permission_ids" json:"permission_ids"`
}

type RoleWithPermissions struct {
	ID                  primitive.ObjectID      `bson:"_id,omitempty" json:"id"`
	Name                string                  `bson:"name" json:"name"`
	Description         string                  `bson:"description" json:"description"`
	PermissionIDs       []primitive.ObjectID    `bson:"permission_ids" json:"permission_ids"`
	PermissionsPopulated []permission.Permission `bson:"permissions_populated,omitempty" json:"permissions_populated,omitempty"`
}


type RoleCreateDto struct {
	Name          string               `bson:"name" json:"name"`
	Description   string               `bson:"description" json:"description"`
	PermissionIDs []primitive.ObjectID `bson:"permission_ids" json:"permission_ids"`
}

type RoleGetDto struct {
	Ids 		[]string 			`json:"ids" validate:"required"`
}