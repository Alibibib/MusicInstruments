package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() error {
	dsn := "host=localhost user=postgres password=qweqwe123 dbname=musicshop port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Подключение к базе данных установлено")

	return nil
}

func GetDB() *gorm.DB {
	return db
}
