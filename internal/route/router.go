package router

import (
	"github.com/gin-gonic/gin"
	Controller "github.com/maanavshah/go-gorm/internal/controller"
	HealthRouter "github.com/maanavshah/go-gorm/internal/route/health"
)

func SetupRouter(r *gin.Engine) {
	HealthRouter.AddHealthRoutes(r)
	Controller.AddBookRoutes(r)
}
