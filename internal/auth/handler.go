package auth

import (
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/response"
)

type AuthHandlerDependency struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, dependency AuthHandlerDependency) {
	handler := &AuthHandler{
		Config: dependency.Config,
	} // Передача конфига
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

// Обработчик для Login
func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
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
		fmt.Println("Register")
	}
}
