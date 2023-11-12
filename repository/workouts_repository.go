package repository

import "SportNotes/models"

type WorkoutsRepository interface {
	Save(workouts models.Workouts)
	// Update(workouts models.Workouts)
	// Delete(workoutsId int)
	// FindById(workoutsId int) (tags models.Workouts, err error)
	// FindAll() []models.Workouts
}
