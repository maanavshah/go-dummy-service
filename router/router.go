package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "pong"})
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "estate server is up and running"})
}

func SetupRouter(r *gin.Engine) {
	r.GET("/health", healthCheck)
	r.GET("/ping", ping)
}
