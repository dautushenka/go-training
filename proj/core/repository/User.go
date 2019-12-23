package repository

import (
	"database/sql"
	"go-training/proj/core/model"
)

type UserRepository interface {
	GetUserById(id int32) (*model.User, error)
	GetUserByLogin(login string) (*model.User, error)
}

type usersRepository struct {
	db *sql.DB
}

func CreateUserRepository(db *sql.DB) UserRepository {
	return &usersRepository{db}
}

func (r *usersRepository) GetUserById(id int32) (*model.User, error) {
	user := model.User{}
	err := r.db.QueryRow("SELECT Id, Login, PasswordHash, FirstName, LastName FROM users WHERE Id=$1", id).Scan(
		&user.Id,
		&user.Login,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
	)
	return &user, err
}

func (r *usersRepository) GetUserByLogin(login string) (*model.User, error) {
	user := model.User{}
	err := r.db.QueryRow("SELECT Id, Login, PasswordHash, FirstName, LastName FROM users WHERE Login=$1", login).Scan(
		&user.Id,
		&user.Login,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
	)
	return &user, err
}
