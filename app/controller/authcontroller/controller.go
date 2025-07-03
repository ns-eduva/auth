package authcontroller

import (
	"eduva-auth/internal/auth"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService auth.UserServiceInterface
}

type UserControllerInterface interface {
	Create(*gin.Context)
}

func NewUserController(userService auth.UserServiceInterface) UserControllerInterface {
	return &userController{
		userService: userService,
	}
}
