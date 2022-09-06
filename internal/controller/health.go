package controller

import (
	Network "github.com/maanavshah/go-dummy-service/common/network"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	Network.ReturnJsonResponse(ctx, http.StatusOK, map[string]string{"msg": "server is up and running"})
}

func HttpPing(ctx *gin.Context) {
	Network.ReturnJsonResponse(ctx, http.StatusOK, map[string]string{"msg": "pong"})
}
