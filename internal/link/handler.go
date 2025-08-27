package link

import (
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
)

type LinkHandlerDeps struct {
	*configs.Config
}

type LinkHandler struct {
	*configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{}
	router.HandleFunc("POST /link", handler.CreateLink())
	router.HandleFunc("PATCH /link/{id}", handler.UpdateLink())
	router.HandleFunc("DELETE /link/{id}", handler.DeleteLink())
	router.HandleFunc("GET /{alias}", handler.GetLink())
}

func (handler *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) GetLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
