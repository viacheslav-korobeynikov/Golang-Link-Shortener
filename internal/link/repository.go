package link

import (
	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/db"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

// Функция создания записи в БД
func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

// Функция получения данных из БД
func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	// Поиск в БД
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

// Функция обновления данных в БД
func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

// Функция удаление данных в БД
func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Функция поиска существующего значения в БД
func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Count() int64 {
	var count int64
	repo.Database.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	return count
}

// Метод получения списка ссылок
func (repo *LinkRepository) GetLinksList(limit, offset int) []Link {
	var links []Link
	repo.Database.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").Limit(limit).
		Offset(offset).
		Scan(&links)
	return links
}
