package user

type User struct {
	ID			string `json:"id"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type UserUseCaseInterface interface {

}

type UserRepoInterface interface {

}