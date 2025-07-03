package authcontroller

import (
	"eduva-auth/internal/auth"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

func (u *userController) Create(c *gin.Context) {
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

	_, err := u.userService.Create(c, userCreateDto)
	if err != nil {
		logger.Ef("LAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		logger.Ef("%v", err)
		ginresponse.InternalServerError(c, err.Error(), err.Error())
		return
	}

	ginresponse.Created(c, "Utilisateur créé avec succès", []string{})
}
