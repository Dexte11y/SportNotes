package requests

import "time"

type UpdateWorkoutsRequest struct {
	IdWorkout int       `json:"id"`
	IdAccount int       `json:"name"`
	Date      time.Time `json:"date"`
}
