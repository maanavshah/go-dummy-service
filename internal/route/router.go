package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	AddHealthRoutes(r)
	apiGroup := r.Group("/api")
	AddBookRoutes(apiGroup)
}
