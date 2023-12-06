package schemas

import (
	"time"
)

type Nutrition struct {
	Id        int       `json:"id,omitempty" db:"id" binding:"-"`
	IdUser    int       `json:"idUser,omitempty" db:"id_user" binding:"-"`
	Type      string    `json:"type" db:"type" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	FoodList  []Food    `json:"foodList"`
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
