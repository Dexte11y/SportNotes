package services

import (
	"SportNotes/data/requests"
	"SportNotes/data/responses"
)

// Сервис аккаутнов
type AccountsService interface {
	// Сервис по созданию аккаунта
	Create(accounts requests.CreateAccountsRequest)

	// Сервис для поиска всех аккаунтов
	FindAll() []responses.AccountsResponse

	// Сервис для поиска аккаунта по Id
	FindById(accountsId int) responses.AccountsResponse

	// Сервис для обновления аккаунта
	// Update(accounts requests.UpdateAccountsRequest)

	// Сервис по удалению аккаунта
	Delete(accountsId int)
}
