package auth

import (
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/req"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/response"
)

type AuthHandlerDependency struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, dependency AuthHandlerDependency) {
	handler := &AuthHandler{
		Config:      dependency.Config,
		AuthService: dependency.AuthService,
	} // Передача конфига
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

// Обработчик для Login
func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		//Возвращаем ответ метода
		data := LoginResponse{
			Token: "123",
		}
		response.Json(w, data, 200)
	}
}

// Обработчик для Register
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		// Вызываем бизнес-логику создания пользователя из service
		handler.AuthService.CreateUser(body.Email, body.Password, body.Name)
	}
}
