package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() *gorm.DB {
	dbURL := "postgres://dev:secret@localhost/activita"

	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB.AutoMigrate(User{}, Activity{})

	return DB
}
