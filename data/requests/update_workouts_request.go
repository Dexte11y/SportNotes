package requests

import "time"

type UpdateWorkoutsRequest struct {
	Id     int       `json:"id"`
	UserId int       `json:"name"`
	Date   time.Time `json:"date"`
}
