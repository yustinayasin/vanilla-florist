package users

import "time"

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserUseCaseInterface interface {
	SignUp(user User) (User, error)
}

type UserRepoInterface interface {
	SignUp(user User) (User, error)
}
