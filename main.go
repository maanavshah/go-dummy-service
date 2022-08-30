package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maanavshah/go-gorm/internal/route"

	Config "github.com/maanavshah/go-gorm/common/config"
)

func main() {
	Config.InitDb(true)
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	router.SetupRouter(app)
	app.Run("0.0.0.0:8000")
}
