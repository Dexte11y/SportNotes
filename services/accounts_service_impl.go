package services

import (
	"SportNotes/data/requests"
	"SportNotes/data/responses"
	"SportNotes/helper"
	"SportNotes/models"
	"SportNotes/repository"

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
func (a *AccountsServiceImpl) Create(account requests.CreateAccountsRequest) {
	err := a.Validate.Struct(account)
	helper.ErrorPanic(err)
	accountModel := models.Account{
		IdAccount:  account.IdAccount,
		Login:      account.Login,
		Name:       account.Name,
		SecondName: account.SecondName,
		Email:      account.Email,
		Password:   account.Password,
	}
	a.AccountsRepository.Save(accountModel)
}

// Поиск всех аккаунтов
func (a *AccountsServiceImpl) FindAll() []responses.AccountsResponse {
	result := a.AccountsRepository.FindAll()

	var accounts []responses.AccountsResponse
	for _, value := range result {
		account := responses.AccountsResponse{
			IdAccount:  value.IdAccount,
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
func (a *AccountsServiceImpl) FindById(accountsId int) responses.AccountsResponse {
	accountData, err := a.AccountsRepository.FindById(accountsId)
	helper.ErrorPanic(err)

	accountResponse := responses.AccountsResponse{
		IdAccount:  accountData.IdAccount,
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
