package permissioncontroller

import (
	"eduva-auth/internal/permission"

	"github.com/gin-gonic/gin"
)

type permissionController struct {
	permissionService permission.PermissionServiceInterface
}

type PermissionControllerInterface interface {
	Create(*gin.Context)
	GetOneById(c *gin.Context)
	UpdateOne(c *gin.Context)
}

func NewPermissionController(permissionService permission.PermissionServiceInterface) PermissionControllerInterface {
	return &permissionController{
		permissionService: permissionService,
	}
}
