package users

import (
	"database/sql"
	users "vanilla-florist/business/user"
)

var db *sql.DB

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database *sql.DB) users.UserRepoInterface {
	//yang direturn adalah interfacenya repo
	return &UserRepository{
		db: database,
	}
}

func (repo *UserRepository) SignUp(user users.User) (users.User, error) {
	userDB := FromUsecase(user)

	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", userDB.Name, userDB.Email, userDB.Password)

	if err != nil {
		return users.User{}, err
	}

	return userDB.ToUsecase(), nil
}
