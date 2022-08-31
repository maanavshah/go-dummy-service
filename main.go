package main

import (
	"github.com/gin-gonic/gin"
	router "github.com/maanavshah/go-dummy-service/internal/route"

	Config "github.com/maanavshah/go-dummy-service/common/config"
)

func main() {
	Config.InitDb()
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	router.SetupRouter(app)
	app.Run("0.0.0.0:8000")
}
