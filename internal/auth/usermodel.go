package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Email     string               `bson:"email" json:"email"`
	Password  string               `bson:"password" json:"password"`
	CrmID     string               `bson:"crm_id" json:"crm_id"`
	RoleIDs   []primitive.ObjectID `bson:"role_ids" json:"role_ids"`
	CreatedAt time.Time            `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time            `bson:"updated_at" json:"updated_at"`
}