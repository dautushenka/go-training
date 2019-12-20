package model

type User struct {
	Id           int32  `json:"id"`
	Login        string `json:"-"`
	PasswordHash string `json:"-"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
