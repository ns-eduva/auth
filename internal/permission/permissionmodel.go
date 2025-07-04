package permission

import "go.mongodb.org/mongo-driver/bson/primitive"

type Permission struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}

type PermissionCreateDto struct {
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
}