package route

import (
	"github.com/gin-gonic/gin"
	HealthController "github.com/maanavshah/go-dummy-service/internal/controller"
)

func AddHealthRoutes(router *gin.Engine) {
	router.GET("/health", HealthController.HealthCheck)
	router.GET("/ping", HealthController.HttpPing)
}
