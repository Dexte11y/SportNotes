package sportnotes

import (
	"errors"
	"time"
)

type Workout struct {
	Id        int       `json:"id" db:"id" binding:"required"`
	IdUser    int       `json:"idUser" db:"id_user" binding:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type UpdWorkout struct {
	Login    *string `json:"login"`
	Name     *string `json:"name"`
	Surname  *string `json:"surname"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (u UpdWorkout) Validate() error {
	if u.Name == nil && u.Surname == nil && u.Login == nil && u.Email == nil && u.Password == nil {
		errors.New("update structure has no values")
	}

	return nil
}
