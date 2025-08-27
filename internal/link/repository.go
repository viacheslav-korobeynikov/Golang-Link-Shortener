package link

import "github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

//Функция создания записи в БД
func (repo *LinkRepository) Crete(link *Link) {

}
