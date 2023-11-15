package requests

import "time"

type CreateWorkoutsRequest struct {
	IdWorkout int       `json:"id"`
	IdAccount int       `json:"name"`
	Date      time.Time `json:"date"`
}
