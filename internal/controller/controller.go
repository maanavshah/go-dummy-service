package controller

import (
	"github.com/gin-gonic/gin"

	BookController "github.com/maanavshah/go-dummy-service/internal/controller/book"
)

func AddBookRoutes(router *gin.Engine) {
	router.POST("/books", BookController.CreateBook)
	router.GET("/books", BookController.ListAllBooks)
}
