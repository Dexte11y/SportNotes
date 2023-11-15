package models

// Модель полей для таблицы users
type Account struct {
	IdAccount  int    `gorm:"type:int"`
	Login      string `gorm:"type:varchar(255)"`
	Name       string `gorm:"type:varchar(255)"`
	SecondName string `gorm:"type:varchar(255)"`
	Email      string `gorm:"type:varchar(255)"`
	Password   string `gorm:"type:varchar(255)"`
}
