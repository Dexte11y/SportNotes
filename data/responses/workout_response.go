package responses

import "time"

type WorkoutsResponse struct {
	IdWorkout int       `json:"id"`
	IdAccount int       `json:"name"`
	Date      time.Time `json:"date"`
}
