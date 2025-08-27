package link

import (
	"fmt"
	"net/http"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/req"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/response"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("POST /link", handler.CreateLink())
	router.HandleFunc("PATCH /link/{id}", handler.UpdateLink())
	router.HandleFunc("DELETE /link/{id}", handler.DeleteLink())
	router.HandleFunc("GET /{hash}", handler.GetLink())
}

func (handler *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Получение body из запроса
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		//Создали сущность в БД
		link := NewLink(body.Url)
		// Записали в репозиторий
		cretedLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Возвращаем ответ
		response.Json(w, cretedLink, 201)
	}
}

func (handler *LinkHandler) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}

func (handler *LinkHandler) GetLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
