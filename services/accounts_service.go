package services

import (
	"SportNotes/schemas"
)

// Сервис аккаутнов
type AccountsService interface {
	// Сервис по созданию аккаунта
	Create(accounts schemas.CreateAccountSchema)

	// Сервис по авторизации аккаунта
	Login(accounts schemas.LoginAccountSchema) bool

	// Сервис для поиска всех аккаунтов
	FindAll() []schemas.ResponseAccountSchema

	// Сервис для поиска аккаунта по Id
	FindById(accountsId int) schemas.ResponseAccountSchema

	// Сервис для обновления аккаунта
	// Update(accounts requests.UpdateAccountsRequest)

	// Сервис по удалению аккаунта
	Delete(accountsId int)
}
