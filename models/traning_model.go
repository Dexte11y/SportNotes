package models

// Модель полей для таблицы training
type Training struct {
	Id          string `gorm:"type:string;primary_key"`
	IdWorkout   int    `gorm:"type:int"`
	Type        string `gorm:"type:string"`
	Name        string `gorm:"type:string"`
	Approaches  string `gorm:"type:string"`
	Repetitions string `gorm:"type:string"`
	Weight      string `gorm:"type:string"`
}
