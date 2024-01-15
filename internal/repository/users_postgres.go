package repository

import (
	"fmt"
	"strings"

	"sportnotes/internal/schemas"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UsersListPostgres struct {
	db *sqlx.DB
}

func NewUsersListPostgres(db *sqlx.DB) *UsersListPostgres {
	return &UsersListPostgres{db: db}
}

func (r *UsersListPostgres) CreateUser(input schemas.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (login, name, surname, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id", usersTable)

	row := r.db.QueryRow(query, input.Login, input.Name, input.Surname, input.Email, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UsersListPostgres) GetAllUsers() ([]schemas.User, error) {
	var users []schemas.User

	query := fmt.Sprintf("SELECT id, login, name, surname, email, password FROM %s", usersTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UsersListPostgres) GetUserByID(id int) (schemas.User, error) {
	var user schemas.User

	query := fmt.Sprintf("SELECT id, login, name, surname, email, password FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UsersListPostgres) UpdateUser(id int, input schemas.UpdUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Login != nil {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argID))
		args = append(args, *input.Login)
		argID++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argID))
		args = append(args, *input.Surname)
		argID++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argID))
		args = append(args, *input.Email)
		argID++
	}

	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argID))
		args = append(args, *input.Password)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", usersTable, setQuery, argID)
	args = append(args, id)

	logrus.Debug("updateQuery: ", query)
	logrus.Debug("args: ", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *UsersListPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, id)

	return err
}
