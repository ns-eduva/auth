package router

import (
	"eduva-auth/docs"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/ginresponse"
	"github.com/nsevenpack/logger/v2/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(r *gin.Engine) {
	pathApiV1 := "api/v1"

	//authService := auth.NewAuthService(initializer.Db)
	//authMiddle := auth.NewAuthMiddleware(authService)
	//authController := authcontroller.NewAuthController(authService)

	docs.SwaggerInfo.BasePath = "/" + pathApiV1
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//v1 := r.Group(pathApiV1)

	//v1Auth := v1.Group("/user")
	// v1Auth.POST("/register", authController.Create)

	r.NoRoute(func(ctx *gin.Context) {
		logger.Wf("Route inconnue : %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		ginresponse.NotFound(ctx, "La route demandée n'existe pas.", "La route demandée n'existe pas.")
		ctx.Abort()
	})
}
