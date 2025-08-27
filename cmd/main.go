package main

import (
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/auth"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/link"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/db"
)

func main() {
	conf := configs.LoadConfig() // Достаем значения конфигов
	_ = db.NewDb(conf)           //Инициализация БД
	router := http.NewServeMux() // Роутниг
	//Обработчики (Handlers)
	//Обработчик авторизации/регистрации
	auth.NewAuthHandler(router, auth.AuthHandlerDependency{
		Config: conf,
	})
	//Обработчик CRUD для ссылок
	link.NewLinkHandler(router, link.LinkHandlerDeps{})
	//Создание сервера
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe() // Подключение к порту
}
