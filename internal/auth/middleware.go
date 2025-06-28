package auth

type authMiddleware struct {
	authService AuthServiceInterface
}

type AuthMiddlewareInterface interface {
}

func NewAuthMiddleware(authService AuthServiceInterface) AuthMiddlewareInterface {
	return &authMiddleware{
		authService: authService,
	}
}