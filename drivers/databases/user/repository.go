package users

import (
	"database/sql"
	"errors"
	"fmt"
	users "vanilla-florist/business/user"
	"vanilla-florist/helpers"
)

type UserRepository struct {
	db *helpers.Database
}

func NewUserRepository(database *helpers.Database) users.UserRepoInterface {
	//yang direturn adalah interfacenya repo
	return &UserRepository{
		db: database,
	}
}

func (repo *UserRepository) SignUp(user users.User) (users.User, error) {
	userDB := FromUsecase(user)

	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return users.User{}, err
	}

	//check if the db is connect
	if db == nil {
		fmt.Println("Database connection is nil")
		return users.User{}, errors.New("database connection is nil")
	}

	_, err = db.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", userDB.Name, userDB.Email, userDB.Password)

	if err != nil {
		fmt.Println("Error in repo: ", err)
		return users.User{}, err
	}

	return userDB.ToUsecase(), nil
}

func (repo *UserRepository) Login(user users.User) (users.User, error) {
	userDB := FromUsecase(user)

	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return users.User{}, err
	}

	if db == nil {
		fmt.Println("Database connection is nil")
		return users.User{}, errors.New("database connection is nil")
	}

	row := db.DB.QueryRow("SELECT * FROM users WHERE email = $1", user.Email)

	//scan jumlah kolom harus sama dan urut
	if err := row.Scan(&userDB.Name, &userDB.Email, &userDB.Password, &userDB.Id); err != nil {
		if err == sql.ErrNoRows {
			return users.User{}, errors.New("User not found")
		}
		return users.User{}, errors.New("Error in database")
	}

	return userDB.ToUsecase(), nil
}
