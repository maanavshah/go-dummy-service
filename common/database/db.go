package database

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	GormLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDatabaseConnection(dbURL string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: GormLogger.Default.LogMode(GormLogger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func InitDb() {
	dbURL := viper.GetString("databaseUrl")
	DB = GetDatabaseConnection(dbURL)
}
