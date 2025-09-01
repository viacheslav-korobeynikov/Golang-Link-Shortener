package link

import (
	"net/http"
	"strconv"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/configs"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/middlware"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/req"
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/response"
	"gorm.io/gorm"
)

type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
	Config         *configs.Config
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.Handle("POST /link", middlware.IsAuthed(handler.CreateLink(), deps.Config))
	router.Handle("PATCH /link/{id}", middlware.IsAuthed(handler.UpdateLink(), deps.Config))
	router.Handle("DELETE /link/{id}", middlware.IsAuthed(handler.DeleteLink(), deps.Config))
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

// Создание ссылки
func (handler *LinkHandler) CreateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Получение body из запроса
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		//Создали сущность в БД (генерим и записываем хэш)
		link := NewLink(body.Url)
		for {
			//Проверяем есть ли такое же значение хэша в БД
			exsitedLink, _ := handler.LinkRepository.GetByHash(link.Hash)
			// Если значение не сушествует - выходим из цикла
			if exsitedLink == nil {
				break
			}
			// Если значение существует, заново генерим хэш
			link.GenerateHash()
		}

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

// Обновление/редактирование ссылки
func (handler *LinkHandler) UpdateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Получение body из запроса
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		//Получение id из path запроса
		idStr := r.PathValue("id")
		// Преобразование строки в число
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Обновление записи
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Возвращение ответа
		response.Json(w, link, 200)
	}
}

// Метод удаления ссылки
func (handler *LinkHandler) DeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Парсим id
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Проверяем наличие записи в БД
		_, err = handler.LinkRepository.GetById(uint(id))
		// Если запись не найдена, то возвращаем ошибку
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Если запись найдена, то удаляем
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Json(w, nil, 200)
	}
}

// Метод получения ссылки и редиректа
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Получение динамического значения из URL
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		//Редирект пользователя по сслыке
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}
