package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json") // Установка хедера
	w.WriteHeader(statusCode)                          // Установка статус-кода ответа
	json.NewEncoder(w).Encode(data)                    //Записываем ответ в json
}
