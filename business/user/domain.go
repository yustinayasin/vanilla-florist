package users

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserUseCaseInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
	FindUser(id int) (User, error)
}

type UserRepoInterface interface {
	SignUp(user User) (User, error)
	Login(user User) (User, error)
	EditUser(user User, id int) (User, error)
	DeleteUser(id int) (User, error)
	FindUser(id int) (User, error)
}
