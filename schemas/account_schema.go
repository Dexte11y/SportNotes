package schemas

type CreateAccountSchema struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	Name       string `json:"name"`
	SecondName string `json:"secondname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginAccountSchema struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UpdateAccountSchema struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type ResponseAccountSchema struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}
