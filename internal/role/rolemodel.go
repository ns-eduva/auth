package role

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name          string               `bson:"name" json:"name"`
	Description   string               `bson:"description" json:"description"`
	PermissionIDs []primitive.ObjectID `bson:"permission_ids" json:"permission_ids"`
}

type RoleCreateDto struct {
	Name          string               `bson:"name" json:"name"`
	Description   string               `bson:"description" json:"description"`
	PermissionIDs []primitive.ObjectID `bson:"permission_ids" json:"permission_ids"`
}