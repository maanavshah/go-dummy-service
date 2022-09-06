package network

import (
	"github.com/gin-gonic/gin"
)

func ReturnJsonResponse(
	ctx *gin.Context,
	httpResponseStatusCode int,
	obj any,
) {
	responseKey := "data"
	if !(299 >= httpResponseStatusCode && httpResponseStatusCode >= 200) {
		responseKey = "error"
	}
	ctx.JSON(httpResponseStatusCode, gin.H{responseKey: obj})
}
