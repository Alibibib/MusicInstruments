package database

import (
	"MusicInstruments/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() error {
	dsn := "host=localhost user=postgres password=qweqwe123 dbname=musicshop port=5432 sslmode=disable"
	var err error
	// Открываем соединение с базой данных PostgreSQL
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Подключение к базе данных установлено")

	// Выполняем миграцию для таблиц User, Category и MusicalInstrument
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.MusicalInstrument{},
	)
	if err != nil {
		log.Println("Ошибка миграции:", err)
		return err
	}
	log.Println("Миграция прошла успешно")

	return nil
}

// Функция для получения объекта базы данных
func GetDB() *gorm.DB {
	return db
}
