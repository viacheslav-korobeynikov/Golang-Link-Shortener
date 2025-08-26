package db

import (
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

// Функция-конструктор БД через горм
func NewDb(conf *configs.Config) *Db {
	//Открытие соединение
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
