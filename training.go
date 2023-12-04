package sportnotes

import "errors"

type Training struct {
	Id          int    `json:"id" db:"id" binding:"required"`
	IdWorkout   int    `json:"idWorkout" db:"id_workout" binding:"required"`
	Name        string `json:"name" db:"name" binding:"required"`
	Approaches  int    `json:"approaches" db:"approaches" binding:"required"`
	Repetitions int    `json:"repetitions" db:"repetitions" binding:"required"`
	Weight      int    `json:"weight" db:"weight" binding:"required"`
}

type TrainingOutput struct {
	Id          int    `json:"idTraining" db:"id" binding:"required"`
	IdWorkout   int    `json:"idWorkout" db:"id_workout" binding:"required"`
	Name        string `json:"name" db:"name" binding:"required"`
	Approaches  int    `json:"approaches" db:"approaches" binding:"required"`
	Repetitions int    `json:"repetitions" db:"repetitions" binding:"required"`
	Weight      int    `json:"weight" db:"weight" binding:"required"`
}

type UpdTraining struct {
	Name        *string `json:"name"`
	Approaches  *int    `json:"approaches"`
	Repetitions *int    `json:"repetitions"`
	Weight      *int    `json:"weight"`
}

func (u UpdTraining) Validate() error {
	if u.Name == nil && u.Approaches == nil && u.Repetitions == nil && u.Weight == nil {
		errors.New("update structure has no values")
	}

	return nil
}
