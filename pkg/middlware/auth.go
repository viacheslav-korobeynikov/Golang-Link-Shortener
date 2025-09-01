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

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		// Создаем новый контекст и добавляем туда ключ и значение для email
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		// Создали новый requst с обогащенным контекстом
		req := r.WithContext(ctx)
		// Передаем новый request
		next.ServeHTTP(w, req)
	})
}
