package config

import (
	"log"

	"github.com/maanavshah/go-dummy-service/internal/model/book"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDatabaseConnection(readOnly bool, dbURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		DryRun: readOnly,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func InitDb() {
	dbURL := "postgres://root:root@127.0.0.1:5432/testv1_golang"
	is_prod_env := viper.GetString("env") == "prod"
	if !is_prod_env {
		db := GetDatabaseConnection(true, dbURL)
		db.AutoMigrate(&model.Book{}) // To show migration SQL only
	}
	DB = GetDatabaseConnection(false, dbURL)
}
