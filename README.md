# Golang-Link-Shortener
Учебный проект приложения по укорачиванию ссылок на Golang


## Разворачивание контейнера с PostgreSQL
```
docker compose up -d
```

## Выполнение автомиграции
```
go run migrations/auto.go
```
## Библиотека для работы с шифрованием
```
go get -u golang.org/x/crypto/bcrypt
```
## Библиотека для работы с JWT
```
go get -u github.com/golang-jwt/jwt/v5
```
