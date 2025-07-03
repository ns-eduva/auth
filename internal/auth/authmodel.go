package auth

import (
	"time"

	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
	RoleIDs  []primitive.ObjectID 	`bson:"role_ids" json:"role_ids"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Ef("Erreur de hashage du mot de passe: %v", err)
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) SetTimeStamps() {
	now := primitive.NewDateTimeFromTime(time.Now())
	if u.CreatedAt == 0 {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
}