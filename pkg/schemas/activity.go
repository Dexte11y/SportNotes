package schemas

import "errors"

type Activity struct {
	ID          int    `json:"id,omitempty" db:"id" binding:"-"`
	IDWorkout   int    `json:"idWorkout,omitempty" db:"id_workout" binding:"-"`
	Name        string `json:"name" db:"name" binding:"required"`
	Approaches  int    `json:"approaches" db:"approaches" binding:"required"`
	Repetitions int    `json:"repetitions" db:"repetitions" binding:"required"`
	Weight      int    `json:"weight" db:"weight" binding:"required"`
}

type UpdActivity struct {
	Name        *string `json:"name"`
	Approaches  *int    `json:"approaches"`
	Repetitions *int    `json:"repetitions"`
	Weight      *int    `json:"weight"`
}

func (u UpdActivity) Validate() error {
	if u.Name == nil && u.Approaches == nil && u.Repetitions == nil && u.Weight == nil {
		errors.New("update structure has no values")
	}

	return nil
}
