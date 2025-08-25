package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
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
		res := LoginResponse{
			Token: "123",
		}
		w.Header().Set("Content-Type", "application/json") // Установка хедера
		w.WriteHeader(201)                                 // Установка статус-кода ответа
		json.NewEncoder(w).Encode(res)                     //Записываем ответ в json
	}
}

// Обработчик для Register
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
