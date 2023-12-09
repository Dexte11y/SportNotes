package repository

import (
	"fmt"
	"time"

	"sportnotes/pkg/schemas"

	"github.com/jmoiron/sqlx"
)

type NutritionsListPostgres struct {
	db *sqlx.DB
}

func NewNutritionListPostgres(db *sqlx.DB) *NutritionsListPostgres {
	return &NutritionsListPostgres{db: db}
}

func (r *NutritionsListPostgres) CreateNutrition(idUser int, input schemas.Nutrition) (int, error) {
	var foodList []schemas.Food
	var idNutrition int
	currentData := time.Now().UTC()

	foodList = append(foodList, input.FoodList...)

	queryNutritions := fmt.Sprintf("INSERT INTO %s (id_user, type, created_at) VALUES ($1, $2, $3) RETURNING id", nutritionsTable)
	row := r.db.QueryRow(queryNutritions, idUser, input.Type, currentData)
	if err := row.Scan(&idNutrition); err != nil {
		return 0, err
	}

	queryFoods := fmt.Sprintf("INSERT INTO %s (id_nutrition, name) VALUES (%d, :name)", foodsTable, idNutrition)
	_, err := r.db.NamedExec(queryFoods, foodList)
	if err != nil {
		return 0, err
	}

	return idNutrition, nil
}

func (r *NutritionsListPostgres) GetNutritionsByParam(id int, startpoint, endpoint string) ([]schemas.Nutrition, error) {
	var nutritionsList []schemas.Nutrition
	var foodsList []schemas.Food

	queryNutritions := fmt.Sprintf("SELECT id, id_user, type, created_at FROM %s WHERE id_user = $1 AND created_at >= $2 AND created_at <= $3", nutritionsTable)
	r.db.Select(&nutritionsList, queryNutritions, id, startpoint, endpoint)

	queryFood := fmt.Sprintf("SELECT id, id_nutrition, name FROM %s", foodsTable)
	r.db.Select(&foodsList, queryFood)

	mergetNutritions := make([]schemas.Nutrition, 0)
	for _, valueNutritions := range nutritionsList {
		for _, valueFoods := range foodsList {
			if valueNutritions.ID == valueFoods.IDNutrition {
				valueNutritions.FoodList = append(valueNutritions.FoodList, valueFoods)
			}
		}
		mergetNutritions = append(mergetNutritions, valueNutritions)
	}

	return mergetNutritions, nil
}

// func (r *NutritionsListPostgres) UpdateNutrition(id int, input schemas.UpdWorkout) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.CreatedAt != nil {
// 		setValues = append(setValues, fmt.Sprintf("login=$%d", argId))
// 		args = append(args, *input.CreatedAt)
// 		argId++
// 	}

// 	setQuery := strings.Join(setValues, ", ")

// 	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d", workoutsTable, setQuery, argId)
// 	args = append(args, id)

// 	logrus.Debug("updateQuery: ", query)
// 	logrus.Debug("args: ", args)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }

func (r *NutritionsListPostgres) DeleteNutrition(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", nutritionsTable)
	_, err := r.db.Exec(query, id)

	return err
}
