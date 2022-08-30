package config

import (
	"fmt"
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
