package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupConfig() {
	viper.SetConfigFile("conf/conf.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func SetupGinAppConfig() *gin.Engine {
	env := viper.GetString("env")
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	return app
}

func GetHostString() string {
	return fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))
}
