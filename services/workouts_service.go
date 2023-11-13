package services

import (
	"SportNotes/data/requests"
	"SportNotes/data/responses"
)

// Сервис тренировок
type WorkoutsService interface {
	// Сервис по созданию тренировки
	Create(workouts requests.CreateWorkoutsRequest)

	// Сервис для поиска всех тренировок
	FindAll() []responses.WorkoutsResponse

	// Сервис для поиска тренировки по Id
	FindById(workoutsId int) responses.WorkoutsResponse

	// Сервис для обновления тренировки
	// Update(workouts requests.UpdateWorkoutsRequest)

	// Сервис по удалению тренировки
	Delete(workoutsId int)
}
