package permissioncontroller

import (
	"eduva-auth/internal/permission"
)

type permissionController struct {
	permissionService permission.PermissionServiceInterface
}

type PermissionControllerInterface interface {
	// Create(*gin.Context)
}

func NewPermissionController(permissionService permission.PermissionServiceInterface) PermissionControllerInterface {
	return &permissionController{
		permissionService: permissionService,
	}
}
