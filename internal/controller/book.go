package controller

import (
	"github.com/gin-gonic/gin"
	database "github.com/maanavshah/go-gorm/internal/db"
	"github.com/maanavshah/go-gorm/internal/model"
	"net/http"
)

type BookPayload struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Desc   string `json:"desc"`
}

func CreateBook(ctx *gin.Context) {
	var payload BookPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := model.Book{Title: payload.Title, Author: payload.Author, Desc: payload.Desc}
	database.DB.Create(&book)
	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func ListBooks(ctx *gin.Context) {
	var books []model.Book
	database.DB.Find(&books)
	ctx.JSON(http.StatusOK, gin.H{"data": books})
}
