package permissioncontroller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (p *permissionController) GetOneById(c *gin.Context) {
	idParam := c.Param("id")

	permissionId, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		logger.Ef("ID invalide : %v", idParam)
		ginresponse.BadRequest(c, "ID invalide", ginresponse.ErrorModel{
			Message: "ID fourni invalide",
			Type:    "Validation",
			Detail:  err.Error(),
		})
		return
	}

	permission, err := p.permissionService.GetOneById(c, permissionId)
	if err != nil {
		if errors.Is(err, errors.New("permission non trouvée")) {
			ginresponse.NotFound(c, "Permission non trouvée", ginresponse.ErrorModel{
				Message: err.Error(),
				Type:    "NotFound",
				Detail:  err.Error(),
			})
			return
		}

		logger.Ef("%v", err)
		ginresponse.InternalServerError(c, "Erreur interne", err.Error())
		return
	}

	ginresponse.Success(c, "Permission récupérée", permission)
}
