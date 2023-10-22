package users

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUseCaseInterface interface {
	SignUp(user User) (User, error)
}

type UserRepoInterface interface {
	SignUp(user User) (User, error)
}
