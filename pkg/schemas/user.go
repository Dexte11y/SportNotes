package schemas

import "errors"

type User struct {
	ID       int    `json:"id,omitempty" db:"id" binding:"-"`
	Login    string `json:"login" db:"login" binding:"required"`
	Name     string `json:"name" db:"name" binding:"required"`
	Surname  string `json:"surname" db:"surname" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type UpdUser struct {
	Login    *string `json:"login"`
	Name     *string `json:"name"`
	Surname  *string `json:"surname"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (u UpdUser) Validate() error {
	if u.Name == nil && u.Surname == nil && u.Login == nil && u.Email == nil && u.Password == nil {
		errors.New("update structure has no values")
	}

	return nil
}
