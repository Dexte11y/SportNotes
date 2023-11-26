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

func (r *UsersListPostgres) GetAll() ([]sportnotes.User, error) {
	var patients []sportnotes.User

	query := fmt.Sprintf("SELECT id, name, surname, birthdate FROM %s", usersTable)
	err := r.db.Select(&patients, query)

	return patients, err
}

func (r *UsersListPostgres) GetById(id int) (sportnotes.User, error) {
	var patient sportnotes.User

	query := fmt.Sprintf("SELECT id, name, surname, birthdate FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&patient, query, id)

	return patient, err
}

func (r *UsersListPostgres) UpdateUser(id int, input sportnotes.UpdUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

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

	if input.BirthDate != nil {
		setValues = append(setValues, fmt.Sprintf("birthdate=$%d", argId))
		args = append(args, *input.BirthDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", usersTable, setQuery, argId)
	args = append(args, id)

	logrus.Debug("updateQuery: %s", query)
	logrus.Debug("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *UsersListPostgres) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", usersTable)
	_, err := r.db.Exec(query, id)

	return err
}
