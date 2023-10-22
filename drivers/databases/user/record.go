package users

import users "vanilla-florist/business/user"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func (user User) ToUsecase() users.User {
	return users.User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUsecaseList(user []User) []users.User {
	var newUsers []users.User

	for _, v := range user {
		newUsers = append(newUsers, v.ToUsecase())
	}
	return newUsers
}

func FromUsecase(user users.User) User {
	return User{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
