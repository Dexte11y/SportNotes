package sportnotes

import (
	"time"
)

type Workout struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	IdUser    int       `json:"idUser" db:"id_user" binding:"required"`
	Type      string    `json:"type" db:"type" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type WorkoutOutputById struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	Type      string    `json:"type" db:"type" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" binding:"required"`
	TrainList []Train
}

type Train struct {
	Id          int    `json:"idTraining" db:"id" binding:"required"`
	Name        string `json:"name" db:"name" binding:"required"`
	Approaches  int    `json:"approaches" db:"approaches" binding:"required"`
	Repetitions int    `json:"repetitions" db:"repetitions" binding:"required"`
	Weight      int    `json:"weight" db:"weight" binding:"required"`
}

// type UpdWorkout struct {
// 	CreatedAt string `json:"createdAt"`
// }

// func (u UpdWorkout) Validate() error {
// 	if u.CreatedAt == nil {
// 		errors.New("update structure has no values")
// 	}

// 	return nil
// }
