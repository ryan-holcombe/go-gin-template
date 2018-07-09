package server

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sythe21/go-gin-template/version"
)

type ApplicationVersion struct {
	version string
	sha     string
	env     string
}

func AppInit(env string) *gin.Engine {

	router := gin.Default()

	router.GET("/internal/health", handleInternalHealthRoute())
	router.GET("/internal/version", handlInternalVersionRoute(env))

	return router
}

func handleInternalHealthRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	}
}

func handlInternalVersionRoute(env string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version":     version.VERSION,
			"gitCommit":   version.GITCOMMIT,
			"environment": env,
		})
	}
}
