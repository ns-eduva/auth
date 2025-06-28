package authcontroller

import (
	"eduva-auth/internal/auth"
)

type authController struct {
	authService auth.AuthServiceInterface
}

type AuthControllerInterface interface {
}

func NewAuthController(authService auth.AuthServiceInterface) AuthControllerInterface {
	return &authController{
		authService: authService,
	}
}
