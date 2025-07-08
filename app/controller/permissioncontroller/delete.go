package permissioncontroller

import (
	"eduva-auth/internal/permission"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (p *permissionController) Delete(c *gin.Context) {
	var idsDto permission.PermissionDeleteDto
	if err := c.ShouldBindJSON(&idsDto); err != nil {
		logger.Ef("Erreur lors de la validation des ids: %s", err.Error())
		ginresponse.BadRequest(c, "Erreur de validation des ids", "Erreur de validation des ids")
		return
	}

	var idsObjectIds []primitive.ObjectID
	for _, v := range idsDto.Ids {
		objectId, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			logger.Ef("Conversion id de la permission impossible: %s", err.Error())
			ginresponse.BadRequest(c, "L'id de la permission est invalide", err.Error())
			return
		}
		idsObjectIds = append(idsObjectIds, objectId)
	}

	if err := p.permissionService.Delete(c, idsObjectIds); err != nil {
		logger.Ef("Erreur lors de la suppression des permissions : %s", err.Error())
		ginresponse.InternalServerError(c, "Impossible de supprimer les permissions", err.Error())
		return
	}

	ginresponse.Success(c, "Les permissions ont été supprimées avec succès", nil)
}
