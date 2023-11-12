package repository

import "SportNotes/models"

// Репозиторий тренировок
type WorkoutsRepository interface {
	// Репозиторий для создания тренировки
	Save(workouts models.Workouts)

	// Репозиторий для поиска всех тренировок
	FindAll() []models.Workouts

	// Репозиторий для поиска тренировки по Id
	FindById(workoutsId int) (tags models.Workouts, err error)

	// Репозиторий для обновление тренировки
	Update(workouts models.Workouts)

	// Репозиторий для удаления тренировки
	// Delete(workoutsId int)
}
