package services

import "SportNotes/data/requests"

type WorkoutsService interface {
	Create(workouts requests.CreateWorkoutsRequest)
	// Update(workouts request.UpdateWorkoutsRequest)
	// Delete(workoutsId int)
	// FindById(workoutsId int) response.WorkoutsResponse
	// FindAll() []response.WorkoutsResponse
}
