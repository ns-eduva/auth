package rolecontroller

import (
	"eduva-auth/internal/role"
)

type roleController struct {
	roleService role.RoleServiceInterface
}

type RoleControllerInterface interface {

}

func NewRoleController(roleService role.RoleServiceInterface) RoleControllerInterface {
	return &roleController{
		roleService: roleService,
	}
}
