package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maanavshah/go-gorm/router"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	router.SetupRouter(app)
	app.Run("0.0.0.0:8000")
}
