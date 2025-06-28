package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Email     string               `bson:"email" json:"email"`
	Password  string               `bson:"password" json:"password"`
	CrmID     string               `bson:"crm_id" json:"crm_id"`
	RoleIDs   []primitive.ObjectID `bson:"role_ids" json:"role_ids"`
	CreatedAt primitive.DateTime   `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime   `bson:"updated_at" json:"updated_at"`
}

type UserCreateDto struct {
	Email    string 				`json:"email" binding:"required,email"`
	Password string 				`json:"password" binding:"required,min=6"`
	Username string 				`json:"username" binding:"required"`
	RoleIDs  []primitive.ObjectID 	`bson:"role_ids" json:"role_ids"`
}