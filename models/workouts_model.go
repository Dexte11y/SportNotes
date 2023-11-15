package models

import "time"

// Модель полей для таблицы workouts
type Workouts struct {
	Id     int       `gorm:"type:int"`
	UserId int       `gorm:"type:int"`
	Date   time.Time `gorm:"type:date"`
}
