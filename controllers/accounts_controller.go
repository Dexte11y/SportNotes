package controllers

import (
	"SportNotes/helper"
	"SportNotes/schemas"
	"SportNotes/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountsController struct {
	accountsService services.AccountsService
}

func NewAccountsController(service services.AccountsService) *AccountsController {
	return &AccountsController{
		accountsService: service,
	}
}

// Контроллер на создание аккаунта
func (controller *AccountsController) Create(c *gin.Context) {
	log.Default().Println("Create accounts")
	createAccountsRequest := schemas.CreateAccountSchema{}
	err := c.BindJSON(&createAccountsRequest)
	helper.ErrorPanic(err)

	controller.accountsService.Create(createAccountsRequest)
	webResponse := schemas.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   createAccountsRequest.Id,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// Контроллер на авторизацию аккаунта
func (controller *AccountsController) Login(c *gin.Context) {
	log.Default().Println("Login accounts")
	loginAccountsRequest := schemas.LoginAccountSchema{}
	err := c.BindJSON(&loginAccountsRequest)
	helper.ErrorPanic(err)

	data := controller.accountsService.Login(loginAccountsRequest)
	if data == true {
		webResponse := schemas.Response{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   "authorized",
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, webResponse)
	} else {
		webResponse := schemas.Response{
			Code:   http.StatusUnauthorized,
			Status: "Bad",
			Data:   "unauthorized",
		}
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusUnauthorized, webResponse)
	}

}

// Контроллер на поиск тренировок
func (controller *AccountsController) FindAll(c *gin.Context) {
	log.Default().Println("FindAll accounts")
	accountsResponse := controller.accountsService.FindAll()
	webResponse := schemas.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   accountsResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// Контроллер на обновление тренировки
// func (controller *AccountsController) Update(c *gin.Context) {
// 	log.Default().Println("Update accounts")
// 	updateAccountsRequest := requests.UpdateAccountsRequest{}
// 	err := c.BindJSON(&updateAccountsRequest)
// 	helper.ErrorPanic(err)

// 	accountId := c.Param("accountId")
// 	id, err := strconv.Atoi(accountId)
// 	helper.ErrorPanic(err)
// 	updateAccountsRequest.Id = id

// 	controller.accountsService.Update(updateAccountsRequest)

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "Ok",
// 		Data:   nil,
// 	}
// 	c.Header("Content-Type", "application/json")
// 	c.JSON(http.StatusOK, webResponse)
// }

// Контроллер на поиск тренировки по Id
func (controller *AccountsController) FindById(c *gin.Context) {
	log.Default().Println("FindById accounts")
	accountId := c.Param("accountId")
	id, err := strconv.Atoi(accountId)
	helper.ErrorPanic(err)

	accountResponse := controller.accountsService.FindById(id)

	webResponse := schemas.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   accountResponse,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}

// Контроллер на удаление тренировки
func (controller *AccountsController) Delete(c *gin.Context) {
	log.Default().Println("Delete accounts")
	accountId := c.Param("accountId")
	id, err := strconv.Atoi(accountId)
	helper.ErrorPanic(err)
	controller.accountsService.Delete(id)

	webResponse := schemas.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, webResponse)
}
