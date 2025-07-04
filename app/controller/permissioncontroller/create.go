package permissioncontroller

import (
	"eduva-auth/internal/permission"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
)

func (p *permissionController) Create(c *gin.Context) {
	var permissionCreateDto permission.PermissionCreateDto
	if err := c.ShouldBindJSON(&permissionCreateDto); err != nil {
		logger.Ef("Erreur de validation: %v", err)
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	_, err := p.permissionService.Create(c, permissionCreateDto)
	if err != nil {
		logger.Ef("%v", err)
		ginresponse.InternalServerError(c, err.Error(), err.Error())
		return
	}

	ginresponse.Created(c, "Permission créé avec succès", []string{})
}
