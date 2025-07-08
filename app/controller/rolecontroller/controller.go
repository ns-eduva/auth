package rolecontroller

import (
	"eduva-auth/internal/role"

	"github.com/gin-gonic/gin"
)

type roleController struct {
	roleService role.RoleServiceInterface
}

type RoleControllerInterface interface {
	Create(c *gin.Context)
	GetByIds(c *gin.Context)
}

func NewRoleController(roleService role.RoleServiceInterface) RoleControllerInterface {
	return &roleController{
		roleService: roleService,
	}
}
