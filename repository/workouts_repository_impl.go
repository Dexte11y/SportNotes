package repository

import (
	"SportNotes/helper"
	"SportNotes/models"

	"gorm.io/gorm"
)

type WorkoutsRepositoryImpl struct {
	Db *gorm.DB
}

func NewWorkoutsRepositoryImpl(Db *gorm.DB) WorkoutsRepository {
	return &WorkoutsRepositoryImpl{Db: Db}
}

// Save implements TagsRepository
func (t *WorkoutsRepositoryImpl) Save(tags models.Workouts) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}
