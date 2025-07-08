package rolecontroller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"eduva-auth/internal/role"
)

func (r *roleController) GetByIds(c *gin.Context) {
	var dto role.RoleGetDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		logger.Ef("Erreur de binding JSON : %v", err)
		ginresponse.BadRequest(c, "JSON invalide", ginresponse.ErrorModel{
			Message: "Le corps de la requête est invalide",
			Type:    "Validation",
			Detail:  err.Error(),
		})
		return
	}

	var objectIDs []primitive.ObjectID
	for _, idStr := range dto.Ids {
		objID, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			logger.Ef("ID invalide : %v", idStr)
			ginresponse.BadRequest(c, "ID invalide", ginresponse.ErrorModel{
				Message: "Un des IDs fournis est invalide",
				Type:    "Validation",
				Detail:  err.Error(),
			})
			return
		}
		objectIDs = append(objectIDs, objID)
	}

	roles, err := r.roleService.GetById(c, objectIDs)
	if err != nil {
		if errors.Is(err, errors.New("aucun role trouvé pour les IDs fournis")) {
			ginresponse.NotFound(c, "Aucun role trouvé", ginresponse.ErrorModel{
				Message: err.Error(),
				Type:    "NotFound",
				Detail:  err.Error(),
			})
			return
		}

		logger.Ef("Erreur lors de la récupération des roles : %v", err)
		ginresponse.InternalServerError(c, "Erreur interne", err.Error())
		return
	}

	ginresponse.Success(c, "Roles récupérés", roles)
}