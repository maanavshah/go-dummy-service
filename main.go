package main

import (
	Config "github.com/maanavshah/go-dummy-service/common/config"
	Router "github.com/maanavshah/go-dummy-service/internal/route"
)

func main() {
	Config.SetupConfig()
	Config.InitDb()
	app := Config.SetupGinAppConfig()
	Router.SetupRouter(app)
	app.Run(Config.GetHostString())
}
