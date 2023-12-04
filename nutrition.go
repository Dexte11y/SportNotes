package sportnotes

import (
	"time"
)

type Nutrition struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	IdUser    int       `json:"idUser" db:"id_user" binding:"required"`
	Type      string    `json:"type" db:"type" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type NutritionOutputByParam struct {
	Id        int          `json:"id" db:"id" binding:"required"`
	IdUser    int          `json:"idUser" db:"id_user" binding:"required"`
	Type      string       `json:"type" db:"type" binding:"required"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at" binding:"required"`
	FoodList  []FoodOutput `json:"foodList"`
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
