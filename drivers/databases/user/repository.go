package users

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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
	if err := row.Scan(&userDB.Id, &userDB.Name, &userDB.Email, &userDB.Password); err != nil {
		if err == sql.ErrNoRows {
			return users.User{}, errors.New("User not found")
		}
		return users.User{}, errors.New("error in database")
	}

	return userDB.ToUsecase(), nil
}

func (repo *UserRepository) EditUser(user users.User, id int) (users.User, error) {
	userDB := FromUsecase(user)

	var newUser User

	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return users.User{}, err
	}

	if db == nil {
		fmt.Println("Database connection is nil")
		return users.User{}, errors.New("database connection is nil")
	}

	_, err = db.DB.Exec("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4", userDB.Name, userDB.Email, userDB.Password, id)

	if err != nil {
		fmt.Println("Error in repo: ", err)
		return users.User{}, err
	}

	err = db.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&newUser.Id, &newUser.Name, &newUser.Email)

	if err != nil {
		log.Fatal(err)
	}

	return newUser.ToUsecase(), nil
}

func (repo *UserRepository) DeleteUser(id int) (users.User, error) {
	var userDb User

	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return users.User{}, err
	}

	if db == nil {
		fmt.Println("Database connection is nil")
		return users.User{}, errors.New("database connection is nil")
	}

	_, err = db.DB.Exec("DELETE FROM users WHERE ID = $1", id)

	//kalo ngecek ga ada id kayak gitu pake result kah?
	if err != nil {
		return users.User{}, err
	}

	return userDb.ToUsecase(), nil
}

func (repo *UserRepository) FindUser(id int) (users.User, error) {
	var newUser User

	//connection database
	db, err := helpers.NewDatabase()

	if err != nil {
		return users.User{}, err
	}

	if db == nil {
		return users.User{}, errors.New("database connection is nil")
	}

	err = db.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&newUser.Id, &newUser.Name, &newUser.Email)

	if err != nil {
		return users.User{}, errors.New("user not found")
	}

	return newUser.ToUsecase(), nil
}
