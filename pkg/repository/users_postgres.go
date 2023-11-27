package repository

import (
	"fmt"
	"strings"

	sportnotes "sportnotes"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UsersListPostgres struct {
	db *sqlx.DB
}

func NewUsersListPostgres(db *sqlx.DB) *UsersListPostgres {
	return &UsersListPostgres{db: db}
}

func (r *UsersListPostgres) CreateUser(input sportnotes.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id, login, name, surname, email, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		usersTable)

	row := r.db.QueryRow(query, input.Id, input.Login, input.Name, input.Surname, input.Email, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UsersListPostgres) GetAllUsers() ([]sportnotes.User, error) {
	var users []sportnotes.User

	query := fmt.Sprintf("SELECT id, login, name, surname, email, password FROM %s", usersTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UsersListPostgres) GetUserById(id int) (sportnotes.User, error) {
	var user sportnotes.User

	query := fmt.Sprintf("SELECT id, login, name, surname, email, password FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UsersListPostgres) UpdateUser(id int, input sportnotes.UpdUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Login != nil {
		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
		args = append(args, *input.Login)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.Password != nil {
		setValues = append(setValues, fmt.Sprintf("password=$%d", argId))
		args = append(args, *input.Password)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", usersTable, setQuery, argId)
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
