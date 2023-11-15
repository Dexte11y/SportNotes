package repository

import (
	// "SportNotes/data/requests"
	"SportNotes/helper"
	"SportNotes/models"
	"errors"

	"gorm.io/gorm"
)

type AccountsRepositoryImpl struct {
	Db *gorm.DB
}

func NewAccountsRepositoryImpl(Db *gorm.DB) AccountsRepository {
	return &AccountsRepositoryImpl{Db: Db}
}

// Сохранение тренировок в БД
func (a *AccountsRepositoryImpl) Save(accounts models.Account) {
	result := a.Db.Create(&accounts)
	helper.ErrorPanic(result.Error)
}

// Поиск всех тренировок из БД
func (a *AccountsRepositoryImpl) FindAll() []models.Account {
	var Accounts []models.Account
	result := a.Db.Find(&Accounts)
	helper.ErrorPanic(result.Error)
	return Accounts
}

// Поиск тренировки по Id в БД
func (a *AccountsRepositoryImpl) FindById(accountsId int) (accounts models.Account, err error) {
	var account models.Account
	result := a.Db.Find(&account, accountsId)
	if result != nil {
		return account, nil
	} else {
		return account, errors.New("account is not found")
	}
}

// Обновление тренировки в БД
// func (a *AccountsRepositoryImpl) Update(accounts models.Account) {
// 	var updateAccounts = requests.UpdateAccountsRequest{
// 		Id:     workouts.Id,
// 		UserId: workouts.UserId,
// 		Date:   workouts.Date,
// 	}
// 	result := a.Db.Model(&accounts).Updates(updateAccounts)
// 	helper.ErrorPanic(result.Error)
// }

// Удаление тренировки в БД
func (a *AccountsRepositoryImpl) Delete(accountsId int) {
	var accounts models.Account
	result := a.Db.Where("id = ?", accounts).Delete(&accounts)
	helper.ErrorPanic(result.Error)
}
