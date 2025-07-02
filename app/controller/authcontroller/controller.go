package authcontroller

import (
	"eduva-auth/internal/auth"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService auth.AuthServiceInterface
}

type AuthControllerInterface interface {
	Create(*gin.Context)
}

func NewAuthController(authService auth.AuthServiceInterface) AuthControllerInterface {
	return &authController{
		authService: authService,
	}
}
