package middlware

import (
	"context"
	"net/http"
	"strings"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/jwt"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем хедер
		authHeader := r.Header.Get("Authorization")
		// Проверяем, что значение содержит префик bearer
		if !strings.HasPrefix(authHeader, "Bearer") {
			// Если префикс отсутствует - возвращаем 401
			writeUnauthed(w)
			return
		}
		// Читаем значение
		token := strings.TrimPrefix(authHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		// Проверяем валидность токена
		if !isValid {
			// Если токен не валидный - возвращаем 401
			writeUnauthed(w)
			return
		}
		// Создаем новый контекст и добавляем туда ключ и значение для email
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		// Создали новый requst с обогащенным контекстом
		req := r.WithContext(ctx)
		// Передаем новый request
		next.ServeHTTP(w, req)
	})
}
