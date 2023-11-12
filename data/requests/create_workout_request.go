package requests

type CreateWorkoutsRequest struct {
	Id   int    `json:"id"`
	Name string `validate:"required,min=1,max=30" json:"name"`
}
