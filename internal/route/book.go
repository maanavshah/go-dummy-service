package route

import (
	"github.com/gin-gonic/gin"
	BookController "github.com/maanavshah/go-dummy-service/internal/controller"
)

func AddBookRoutes(router *gin.RouterGroup) {
	router.POST("/books", BookController.CreateBook)
	router.GET("/books", BookController.ListAllBooks)
}
