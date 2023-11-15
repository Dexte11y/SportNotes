package repository

import "SportNotes/models"

// Репозиторий тренировок
type WorkoutsRepository interface {
	// Репозиторий для создания тренировки
	Save(workouts models.Workout)

	// Репозиторий для поиска всех тренировок
	FindAll() []models.Workout

	// Репозиторий для поиска тренировки по Id
	FindById(workoutsId int) (workouts models.Workout, err error)

	// Репозиторий для обновление тренировки
	// Update(workouts models.Workout)

	// Репозиторий для удаления тренировки
	Delete(workoutsId int)
}
