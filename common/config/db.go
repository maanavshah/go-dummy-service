package config

import (
	"log"

	"github.com/maanavshah/go-gorm/internal/model/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb(autoMigrate bool) {
	dbURL := "postgres://root:root@localhost:5432/testv1_golang"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if autoMigrate {
		db.AutoMigrate(&model.Book{})
	}
	DB = db
}
