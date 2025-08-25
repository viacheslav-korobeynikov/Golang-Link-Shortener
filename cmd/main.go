package main

import (
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/auth"
)

func main() {
	conf := configs.LoadConfig() // Достаем значения конфигов
	router := http.NewServeMux() // Роутниг
	auth.NewAuthHandler(router, auth.AuthHandlerDependency{
		Config: conf,
	})
	//Создание сервера
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe() // Подключение к порту
}
