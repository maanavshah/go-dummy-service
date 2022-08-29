package db

import (
	"log"

	"github.com/maanavshah/go-gorm/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbURL := "postgres://maanav:root@localhost:5432/library_db"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.Book{})

	DB = db
}
