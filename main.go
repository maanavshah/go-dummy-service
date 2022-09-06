package main

import (
	Config "github.com/maanavshah/go-dummy-service/common/config"
	Database "github.com/maanavshah/go-dummy-service/common/database"
	Router "github.com/maanavshah/go-dummy-service/internal/route"
)

func main() {
	Config.SetupConfig()
	Database.InitDb()
	app := Config.SetupGinAppConfig()
	Router.SetupRouter(app)
	app.Run(Config.GetHostString())
}
