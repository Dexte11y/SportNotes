package repository

import (
	"fmt"

	"sportnotes/pkg/schemas"

	"github.com/jmoiron/sqlx"
)

type FoodsListPostgres struct {
	db *sqlx.DB
}

func NewFoodsListPostgres(db *sqlx.DB) *FoodsListPostgres {
	return &FoodsListPostgres{db: db}
}

func (r *FoodsListPostgres) CreateFood(input schemas.Food) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (id, id_nutrition, name) VALUES ($1, $2, $3) RETURNING  id ", foodsTable)
	row := r.db.QueryRow(query, input.Id, input.IdNutrition, input.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *FoodsListPostgres) GetAllFoods() ([]schemas.Food, error) {
	var foods []schemas.Food

	query := fmt.Sprintf("SELECT id, id_nutrition, name FROM %s", foodsTable)
	err := r.db.Select(&foods, query)

	return foods, err
}

func (r *FoodsListPostgres) GetFoodById(id int) (schemas.Food, error) {
	var food schemas.Food

	query := fmt.Sprintf("SELECT id, id_nutrition, name FROM %s WHERE id = $1", foodsTable)
	err := r.db.Get(&food, query, id)

	return food, err
}

// func (r *FoodsListPostgres) UpdateFood(id int, input schemas.UpdFood) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.Name != nil {
// 		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
// 		args = append(args, *input.Name)
// 		argId++
// 	}

// 	setQuery := strings.Join(setValues, ", ")

// 	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", trainingsTable, setQuery, argId)
// 	args = append(args, id)

// 	logrus.Debug("updateQuery: ", query)
// 	logrus.Debug("args: ", args)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }

func (r *FoodsListPostgres) DeleteFood(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", foodsTable)
	_, err := r.db.Exec(query, id)

	return err
}
