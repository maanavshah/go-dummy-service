package health

import (
	"github.com/gin-gonic/gin"
	Network "github.com/maanavshah/go-dummy-service/common/network"
	"net/http"
)

func AddHealthRoutes(router *gin.Engine) {
	router.GET("/health", HealthCheck)
	router.GET("/ping", HttpPing)
}

func HealthCheck(ctx *gin.Context) {
	Network.ReturnJsonResponse(ctx, http.StatusOK, map[string]string{"msg": "estate server is up and running"})
}

func HttpPing(ctx *gin.Context) {
	Network.ReturnJsonResponse(ctx, http.StatusOK, map[string]string{"msg": "pong"})
}
