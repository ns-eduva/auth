package auth

type userMiddleware struct {
	userService UserServiceInterface
}

type UserMiddlewareInterface interface {
}

func NewUserMiddleware(userService UserServiceInterface) UserMiddlewareInterface {
	return &userMiddleware{
		userService: userService,
	}
}