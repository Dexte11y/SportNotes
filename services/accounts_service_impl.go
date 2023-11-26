package services

import (
	"SportNotes/helper"
	"SportNotes/models"
	"SportNotes/repository"
	"SportNotes/schemas"

	"github.com/go-playground/validator"
)

type AccountsServiceImpl struct {
	AccountsRepository repository.AccountsRepository
	Validate           *validator.Validate
}

func NewAccountsServiceImpl(AccountRepository repository.AccountsRepository, validate *validator.Validate) AccountsService {
	return &AccountsServiceImpl{
		AccountsRepository: AccountRepository,
		Validate:           validate,
	}
}

// Реализация сервиса аккаунтов
// Создание аккаунта
func (a *AccountsServiceImpl) Create(account schemas.CreateAccountSchema) {
	err := a.Validate.Struct(account)
	helper.ErrorPanic(err)
	accountModel := models.Account{
		Id:         account.Id,
		Login:      account.Login,
		Name:       account.Name,
		SecondName: account.SecondName,
		Email:      account.Email,
		Password:   account.Password,
	}
	a.AccountsRepository.Save(accountModel)
}

// Авторизация аккаунта

func (a *AccountsServiceImpl) Login(account schemas.LoginAccountSchema) bool {
	err := a.Validate.Struct(account)
	helper.ErrorPanic(err)
	accountModel := models.Account{
		Login:    account.Login,
		Password: account.Password,
	}
	return a.AccountsRepository.Login(accountModel)
}

// Поиск всех аккаунтов
func (a *AccountsServiceImpl) FindAll() []schemas.ResponseAccountSchema {
	result := a.AccountsRepository.FindAll()

	var accounts []schemas.ResponseAccountSchema
	for _, value := range result {
		account := schemas.ResponseAccountSchema{
			Id:         value.Id,
			Login:      value.Login,
			Name:       value.Name,
			SecondName: value.SecondName,
			Email:      value.Email,
			Password:   value.Password,
		}
		accounts = append(accounts, account)
	}

	return accounts
}

// Обновление аккаунта
// func (a *AccountsServiceImpl) Update(accounts requests.UpdateAccountsRequest) {
// 	accountsData, err := a.AccountsRepository.FindById(accounts.Id)
// 	helper.ErrorPanic(err)
// 	accountData.Date = accounts.Date
// 	a.AccountsRepository.Update(accountData)
// }

// Поиск аккаунта по Id
func (a *AccountsServiceImpl) FindById(accountsId int) schemas.ResponseAccountSchema {
	accountData, err := a.AccountsRepository.FindById(accountsId)
	helper.ErrorPanic(err)

	accountResponse := schemas.ResponseAccountSchema{
		Id:         accountData.Id,
		Login:      accountData.Login,
		Name:       accountData.Name,
		SecondName: accountData.SecondName,
		Email:      accountData.Email,
		Password:   accountData.Password,
	}
	return accountResponse
}

// Удаление аккаунта
func (a *AccountsServiceImpl) Delete(accountsId int) {
	a.AccountsRepository.Delete(accountsId)
}
