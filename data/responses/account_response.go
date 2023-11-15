package responses

type AccountsResponse struct {
	IdAccount  int    `json:"idAccount"`
	Login      string `json:"login"`
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
