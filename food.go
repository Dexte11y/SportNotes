package sportnotes

import (
	"time"
)

type Food struct {
	Id          int    `json:"id" db:"id" binding:"required"`
	IdNutrition int    `json:"idNutrition" db:"id_nutrition" binding:"required"`
	Name        string `json:"name" db:"name" binding:"required"`
}

type FoodOutput struct {
	Id          int    `json:"idFood" db:"id" binding:"required"`
	IdNutrition int    `json:"idNutrition" db:"id_nutrition" binding:"required"`
	Name        string `json:"name" db:"name" binding:"required"`
}

type UpdFood struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	IdUser    int       `json:"idUser" db:"id_user" binding:"required"`
	Type      string    `json:"type" db:"type" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" binding:"required"`
}

// func (u UpdWorkout) Validate() error {
// 	if u.CreatedAt == nil {
// 		errors.New("update structure has no values")
// 	}

// 	return nil
// }
