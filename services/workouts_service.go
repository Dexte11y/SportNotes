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

	// Сервис для обновления тренировки
	// Update(workouts request.UpdateWorkoutsRequest)

	// Сервис по удалению тренировки
	// Delete(workoutsId int)

	// Сервис для поиска тренировки по Id
	// FindById(workoutsId int) response.WorkoutsResponse
}
