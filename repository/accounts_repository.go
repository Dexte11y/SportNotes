package repository

import "SportNotes/models"

// Репозиторий аккаунтов
type AccountsRepository interface {
	// Репозиторий для создания аккаунта
	Save(accounts models.Account)

	// Репозиторий для авторизации аккаунта
	Login(accounts models.Account) bool

	// Репозиторий для поиска всех аккаунтов
	FindAll() []models.Account

	// Репозиторий для поиска аккаунта по Id
	FindById(accountsId int) (accounts models.Account, err error)

	// Репозиторий для обновление аккаунта
	// Update(accounts models.Account)

	// Репозиторий для удаления аккаунта
	Delete(accountsId int)
}
