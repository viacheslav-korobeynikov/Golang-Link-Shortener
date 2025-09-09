package main

import (
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/auth"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/link"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/stat"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/user"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/db"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/middlware"
)

func main() {
	conf := configs.LoadConfig() // Достаем значения конфигов
	db := db.NewDb(conf)         //Инициализация БД
	router := http.NewServeMux() // Роутниг
	//Репозитории
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)

	//Обработчики (Handlers)
	//Обработчик авторизации/регистрации
	auth.NewAuthHandler(router, auth.AuthHandlerDependency{
		Config:      conf,
		AuthService: authService,
	})
	//Обработчик CRUD для ссылок
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		StatRepository: statRepository,
		Config:         conf,
	})

	// Middlewares
	stack := middlware.Chain(
		middlware.CORS,
		middlware.Logging,
	)

	//Создание сервера
	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router), // Добавлен chain для middleware
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe() // Подключение к порту
}
