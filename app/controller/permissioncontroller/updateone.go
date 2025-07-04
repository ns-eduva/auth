package permissioncontroller

import (
	"eduva-auth/internal/permission"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *permissionController) UpdateOne(c *gin.Context) {
	var permissionUpdateDto permission.PermissionUpdateDto
	if err := c.ShouldBindJSON(&permissionUpdateDto); err != nil {
		logger.Ef("Erreur de validation : %s", err.Error())
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	id := c.Param("id")
	permissionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Ef("impossible de modifier la permission : %s", err.Error())
		ginresponse.BadRequest(c, "impossible de modifier la permission", err.Error())
		return
	}

	if err := p.permissionService.UpdateOne(c, permissionId, permissionUpdateDto); err != nil {
		logger.Ef("Erreur lors de la mise à jour de la permission : %s", err.Error())
		ginresponse.InternalServerError(c, "Erreur lors de la mise à jour de la permission", err.Error())
		return
	}

	ginresponse.Success(c, "Permission mise à jour avec succès", nil)
}
