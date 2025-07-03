package authcontroller

import (
	"eduva-auth/internal/auth"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *authController) Create(c *gin.Context) {
	var userCreateDto auth.UserCreateDto
	if err := c.ShouldBindJSON(&userCreateDto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	if userCreateDto.Email == "" || userCreateDto.Password == "" {
		ginresponse.BadRequest(c, "Email et mot de passe sont requis", ginresponse.ErrorModel{
			Message: "Email et mot de passe sont requis",
			Type:    "Validation",
			Detail:  "Champs manquants",
		})
		return
	}

	newUser := &auth.User{
		Email:    userCreateDto.Email,
		Password: userCreateDto.Password,
		RoleIDs:  userCreateDto.RoleIDs,
	}

	if newUser.RoleIDs == nil {
		newUser.RoleIDs = []primitive.ObjectID{}
	}

	err := a.authService.Create(c, newUser)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.InternalServerError(c, err.Error(), err.Error())
		return
	}

	ginresponse.Created(c, "Utilisateur créé avec succès", []string{})
}
