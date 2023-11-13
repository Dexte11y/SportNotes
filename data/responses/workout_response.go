package responses

import "time"

type WorkoutsResponse struct {
	Id     int       `json:"id"`
	UserId int       `json:"name"`
	Date   time.Time `json:"date"`
}
