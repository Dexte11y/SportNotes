package requests

type UpdateAccountsRequest struct {
	IdAccount  int    `json:"idAccount"`
	Login      string `json:"login"`
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
