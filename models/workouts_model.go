package models

type Workouts struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}
