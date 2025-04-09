package database

import (
	"MusicInstruments/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() error {
	dsn := "host=localhost user=postgres password=qweqwe123 dbname=musicshop port=5432 sslmode=disable" // поменял dbname на musicshop
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Подключение к базе данных успешно установлено")

	err = db.Table("public.users").AutoMigrate(&models.User{})
	if err != nil {
		log.Println("Ошибка при миграции:", err)
		return err
	}

	log.Println("Миграция прошла успешно")
	return nil
}
func GetDB() *gorm.DB {
	return db
}
