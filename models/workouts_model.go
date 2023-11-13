package models

// Модель полей для таблицы workouts
type Workouts struct {
	Id   int    `gorm:"type:int"`
	Name string `gorm:"type:varchar(255)"`
}

// Модель полей для таблицы users
type Users struct {
	Id       int    `gorm:"type:int"`
	Name     string `gorm:"type:varchar(255)"`
	LastName string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
