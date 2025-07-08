package rolecontroller

import (
	"eduva-auth/internal/role"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func (r *roleController) UpdateOne(c *gin.Context) {
	var roleUpdateDto role.RoleUpdateDto
	if err := c.ShouldBindJSON(&roleUpdateDto); err != nil {
		logger.Ef("Erreur de validation : %s", err.Error())
		ginresponse.BadRequest(c, "Erreur de validation", ginresponse.ErrorModel{
			Message: err.Error(),
			Type:    "Validation",
			Detail:  fmt.Sprintf("%v", err),
		})
		return
	}

	id := c.Param("id")
	roleId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Ef("impossible de modifier la role : %s", err.Error())
		ginresponse.BadRequest(c, "impossible de modifier la role", err.Error())
		return
	}

	if err := r.roleService.UpdateOne(c, roleId, roleUpdateDto); err != nil {
		logger.Ef("Erreur lors de la mise à jour de la role : %s", err.Error())
		ginresponse.InternalServerError(c, "Erreur lors de la mise à jour de la role", err.Error())
		return
	}

	ginresponse.Success(c, "Role mise à jour avec succès", nil)
}
