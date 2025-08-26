package link

import (
	"math/rand"

	"gorm.io/gorm"
)

// Описание модели
type Link struct {
	gorm.Model        // Системные поля: ID, created_at, updated_at, deleted_at
	Url        string `json:"url"`
	Hash       string `json:"hash" gorm:"uniqueIndex"`
}

// Функция-конструктор
func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: RandStringRunes(10),
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMOPQRSTUVWXYZ")

// Функция генерации хэша
func RandStringRunes(n int) string {
	sliceRunes := make([]rune, n)
	for i := range sliceRunes {
		sliceRunes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(sliceRunes)
}
