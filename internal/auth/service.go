package auth

import (
	"errors"

	"github.com/viacheslav-korobeynikov/Golang-Link-Shortener/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

// Бизнес логика создания пользователя
func (service *AuthService) CreateUser(email, password, name string) (string, error) {
	// Проверяем существует ли пользователь с таким email в БД
	existedUser, _ := service.UserRepository.FindUserByEmail(email)
	// Если существует возвращаем ошибку
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	//Шифруем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Формируем данные для созадния
	user := &user.User{
		Email:    email,
		Password: string(hashedPassword), // TODO: зашифровать пароль
		Name:     name,
	}
	// Создаем пользователя в БД
	_, err = service.UserRepository.Create(user)
	// Если создать не удалось - возвращаем ошибку
	if err != nil {
		return "", err
	}
	// Если удалось создать - возвращаем email
	return user.Email, nil
}
