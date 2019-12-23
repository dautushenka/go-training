package fixtures

import "go-training/proj/core/model"

var Users = []model.User{
	{
		Id:        1,
		Login:     "user",
		FirstName: "Alex",
		LastName:  "Petrov",
	},
	{
		Id:        2,
		Login:     "admin",
		FirstName: "Denis",
		LastName:  "Autushenka",
	},
}
