package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	Config "github.com/maanavshah/go-gorm/common/config"
	Network "github.com/maanavshah/go-gorm/common/network"
	BookDTO "github.com/maanavshah/go-gorm/internal/dto/book"
	"github.com/maanavshah/go-gorm/internal/model/book"
)

func CreateBook(ctx *gin.Context) {
	var payload BookDTO.CreateBookRequestDTO
	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		Network.ReturnJsonResponse(ctx, http.StatusBadRequest, err.Error())
	} else {
		book := model.Book{Title: payload.Title, Author: payload.Author, Desc: payload.Desc}
		Config.DB.Create(&book)
		Network.ReturnJsonResponse(ctx, http.StatusOK, book)
	}
}

func ListAllBooks(ctx *gin.Context) {
	var books []model.Book
	Config.DB.Find(&books)
	Network.ReturnJsonResponse(ctx, http.StatusOK, books)
}
