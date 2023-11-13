package requests

import "time"

type CreateWorkoutsRequest struct {
	Id     int       `json:"id"`
	UserId int       `json:"name"`
	Date   time.Time `json:"date"`
}
