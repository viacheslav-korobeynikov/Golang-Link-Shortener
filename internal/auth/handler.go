package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

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
	return func(w http.ResponseWriter, request *http.Request) {
		//Чтение body
		var payload LoginRequest
		err := json.NewDecoder(request.Body).Decode(&payload)
		if err != nil {
			response.Json(w, err.Error(), 400)
			return
		}
		// Простая валидация, что передана не пустая строка
		if payload.Email == "" {
			response.Json(w, "Email required", 400)
			return
		}
		//Проверка регулярным выражением
		reg, _ := regexp.Compile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`)
		if !reg.MatchString(payload.Email) {
			response.Json(w, "Wrong email", 400)
		}
		if payload.Password == "" {
			response.Json(w, "Password required", 400)
			return
		}
		fmt.Println(payload)
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
