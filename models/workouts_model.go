package models

// Модель полей для таблицы workouts
type Workouts struct {
	Id   int    `gorm:"type:int"`
	Name string `gorm:"type:varchar(255)"`
}
