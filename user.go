package sportnotes

import "errors"

// type User struct {
// 	Id        int    `json:"-"`
// 	Name      string `json:"name" db:"name" binding:"required" example:"Mihail"`
// 	Surname   string `json:"surname" db:"surname" binding:"required" example:"Kravcov"`
// 	BirthDate string `json:"birthdate" binding:"required" time_format:"2006-01-25" example:"2006-01-25"`
// }

type User struct {
	Id       int    `json:"id" db:"id" binding:"required"`
	Login    string `json:"login" db:"login" binding:"required"`
	Name     string `json:"name" db:"name" binding:"required"`
	Surname  string `json:"surname" db:"surname" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type UpdUser struct {
	Name      *string `json:"name" example:"Igor"`
	Surname   *string `json:"surname" example:"Vasilev"`
	BirthDate *string `json:"birthdate" time_format:"2006-01-02" example:"2000-05-13"`
}

func (u UpdUser) Validate() error {
	if u.Name == nil && u.Surname == nil && u.BirthDate == nil {
		errors.New("update structure has no values")
	}

	return nil
}
