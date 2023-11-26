package schemas

import "time"

type CreateWorkoutSchema struct {
	Id        int `json:"id"`
	IdAccount int `json:"id_account"`
	// CreatedAt time.Time `json:"date"`
}

type UpdateWorkoutSchema struct {
	Id        int `json:"id"`
	IdAccount int `json:"id_account"`
	// Date      time.Time `json:"date"`
}

type ResponseWorkoutSchema struct {
	Id        int       `json:"id"`
	IdAccount int       `json:"id_account"`
	CreatedAt time.Time `json:"date"`
}
