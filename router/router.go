package router

import (
	"github.com/maanavshah/go-gorm/internal/controller"
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
	r.POST("/books", controller.CreateBook)
	r.GET("/books", controller.ListBooks)
}
