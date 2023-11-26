package services

import (
	"SportNotes/schemas"
)

// Сервис тренировок
type WorkoutsService interface {
	// Сервис по созданию тренировки
	Create(workouts schemas.CreateWorkoutSchema)

	// Сервис для поиска всех тренировок
	FindAll() []schemas.ResponseWorkoutSchema

	// Сервис для поиска тренировки по Id
	FindById(workoutsId int) schemas.ResponseWorkoutSchema

	// Сервис для обновления тренировки
	// Update(workouts requests.UpdateWorkoutsRequest)

	// Сервис по удалению тренировки
	Delete(workoutsId int)
}
