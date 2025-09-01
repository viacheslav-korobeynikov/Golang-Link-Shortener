package auth

import (
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/jwt"
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
		// Авторизация пользователя
		email, err := handler.AuthService.LoginUser(body.Email, body.Password)
		// Если авторизация завершиалсь ошибкой - возвращаем 401
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// Если авторизация завершилась успешно - генерируем access_token
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		// Если не удалось сгенерировать токен - возвращаем 500
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//Возвращаем ответ метода
		data := LoginResponse{
			Token: token,
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
		email, err := handler.AuthService.CreateUser(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{
			Email: email,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegisterResponse{
			Token: token,
		}
		response.Json(w, data, 200)
	}
}
