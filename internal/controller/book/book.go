package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Database "github.com/maanavshah/go-dummy-service/common/database"
	Network "github.com/maanavshah/go-dummy-service/common/network"
	BookDTO "github.com/maanavshah/go-dummy-service/internal/dto/book"
	model "github.com/maanavshah/go-dummy-service/internal/model/book"
)

func CreateBook(ctx *gin.Context) {
	var payload BookDTO.CreateBookRequestDTO
	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		Network.ReturnJsonResponse(ctx, http.StatusBadRequest, err.Error())
	} else {
		book := model.Book{Title: payload.Title, Author: payload.Author, Desc: payload.Desc}
		Database.DB.Create(&book)
		Network.ReturnJsonResponse(ctx, http.StatusOK, book)
	}
}

func ListAllBooks(ctx *gin.Context) {
	var books []model.Book
	Database.DB.Find(&books)
	Network.ReturnJsonResponse(ctx, http.StatusOK, books)
}
