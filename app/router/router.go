package router

import (
	"eduva-auth/app/controller/authcontroller"
	"eduva-auth/app/controller/permissioncontroller"
	"eduva-auth/docs"
	"eduva-auth/internal/auth"
	"eduva-auth/internal/permission"

	initializer "eduva-auth/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine) {
	pathApiV1 := "api/v1"

	userService := auth.NewUserService(initializer.Db)
	//authMiddle := auth.NewAuthMiddleware(userService)
	userController := authcontroller.NewUserController(userService)

	permissionService := permission.NewPermissionService(initializer.Db)
	permissionController := permissioncontroller.NewPermissionController(permissionService)

	docs.SwaggerInfo.BasePath = "/" + pathApiV1
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group(pathApiV1)

	v1Auth := v1.Group("/auth")
	v1Auth.POST("/register", userController.Create)

	v1Permission := v1.Group("/permission")
	v1Permission.POST("/create", permissionController.Create)
	v1Permission.GET("/:id", permissionController.GetOneById)
	v1Permission.PUT("/:id", permissionController.UpdateOne)

	r.NoRoute(func(ctx *gin.Context) {
		logger.Wf("Route inconnue : %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		ginresponse.NotFound(ctx, "La route demandée n'existe pas.", "La route demandée n'existe pas.")
		ctx.Abort()
	})
}
