package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/link"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	//Открытие соединение
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Вызываем автомиграцию
	db.AutoMigrate(&link.Link{})
}
