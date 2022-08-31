package router

import (
	"github.com/gin-gonic/gin"
	Controller "github.com/maanavshah/go-dummy-service/internal/controller"
	HealthRouter "github.com/maanavshah/go-dummy-service/internal/route/health"
)

func SetupRouter(r *gin.Engine) {
	HealthRouter.AddHealthRoutes(r)
	Controller.AddBookRoutes(r)
}
