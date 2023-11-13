package models

import "time"

// Модель полей для таблицы workouts
type Workouts struct {
	Id     int       `gorm:"type:int"`
	UserId int       `gorm:"type:int"`
	Date   time.Time `gorm:"type:date"`
}

// Модель полей для таблицы users
type Users struct {
	Id       int    `gorm:"type:int"`
	Name     string `gorm:"type:varchar(255)"`
	LastName string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
