package authcontroller

import (
	"eduva-auth/internal/auth"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

func (a *authController) Create(c *gin.Context) {
	var authCreateDto auth.AuthCreateDto
	if err := c.ShouldBindJSON(&authCreateDto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	newAuth := &auth.Auth{
		Email:    authCreateDto.Email,
		Password: authCreateDto.Password,
		RoleIDs:  authCreateDto.RoleIDs,
	}

	err := a.authService.Create(c, newAuth)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.InternalServerError(c, err.Error(), err.Error())
		return
	}

	ginresponse.Created(c, "Auth créé avec succès", []string{})
}
