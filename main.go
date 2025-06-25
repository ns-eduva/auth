package main

import (
	"eduva-auth/docs"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nsevenpack/env/env"
	"github.com/nsevenpack/logger/v2/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	logger.Init(env.Get("APP_ENV"))
}

// @title Eduva Auth
// @version 1.0
// @description API service Auth d'Eduva
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in headers
// @name Authorization
func main() {
	s := gin.Default()
	host := "0.0.0.0"
	hostTraefik := extractStringInBacktick(env.Get("HOST_TRAEFIK"))
	port := env.Get("PORT")

	setSwaggerOpt(hostTraefik)
	s.GET("/api-doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	infoServer(hostTraefik)
	if err := s.Run(host + ":" + port); err != nil {
		logger.Ef("Une erreur est survenue au lancement du serveur : %v", err)
	}
}

func infoServer(hostTraefik string) {
	logger.If("Lancement du serveur : https://%v", hostTraefik)
	logger.If("Lancement du Swagger : https://%v/api-doc/index.html", hostTraefik)
}

func extractStringInBacktick(s string) string {
	start := strings.Index(s, "`")
	end := strings.LastIndex(s, "`")

	if start == -1 || end == -1 || start == end {
		return ""
	}

	return s[start+1 : end]
}

func setSwaggerOpt(hostTraefik string) {
	docs.SwaggerInfo.Host = hostTraefik
}