package user

import "github.com/viacheslav-korobeynikov/Golang-Link-Shortener/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

// Функция создания пользователя в БД
func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

//Поиск пользователя по email
func (repo *UserRepository) FindUserByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
