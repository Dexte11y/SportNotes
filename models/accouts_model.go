package models

// Модель полей для таблицы account
type Account struct {
	Id         int    `gorm:"type:int;primary_key"`
	Login      string `gorm:"type:string"`
	Name       string `gorm:"type:string"`
	SecondName string `gorm:"type:string"`
	Email      string `gorm:"type:string"`
	Password   string `gorm:"type:string"`
}
