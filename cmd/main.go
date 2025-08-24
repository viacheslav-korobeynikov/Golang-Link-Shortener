package main

import (
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/auth"
)

func main() {
	router := http.NewServeMux() // Роутниг
	auth.NewAuthHandler(router)
	//Создание сервера
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe() // Подключение к порту
}
