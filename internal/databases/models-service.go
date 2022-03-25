// Package databases - реализует весь функционал необходимый для взаимодействия с базами данных
package databases

import (
	"paintings-shop/packages/authentication"

	uuid "github.com/gofrs/uuid"
)

// User - тип для хранения данных о пользователе в базе данных
type User struct {
	GUID         uuid.UUID
	Role         string
	Email        string
	Phone        string
	Name         string
	IsAdmin      bool
	Confirmed    bool
	Disabled     bool
	SecondFactor bool
	Lang         string
}

// Users - тип для хранения списка пользователей
type Users []User

// UsersResponse  - тип для возврата с ответом,
// описывающий список пользователей для постраничной разбивки
type UsersResponse struct {
	Users  Users
	Total  int
	Offset int
	Limit  int
}

// TOTPSecret - секрет для Time Based One Time Password
type TOTPSecret struct {
	UserID    uuid.UUID
	Secret    string
	EncKey    []byte
	Confirmed bool
}

// TOTPResponse - тип для подтверждения наличия секрета
type TOTPResponse struct {
	UserID    uuid.UUID
	Confirmed bool
}

//
// Типы для упрощения создания таблиц
//

// NamedCreateStatement - тип для хранения имени таблицы и кода для её создания в базе
type NamedCreateStatement struct {
	TableName       string
	CreateStatement string
}

// NamedCreateStatements - массив объектов с названием таблицы и кодом для её создания
type NamedCreateStatements []NamedCreateStatement

// SessionsResponse - структура возвращаемая в ответ на запрос сессий
type SessionsResponse struct {
	Sessions
	Total  int
	Offset int
	Limit  int
}

// Sessions - структура описывающая список активных сессий
type Sessions []authentication.ActiveToken
