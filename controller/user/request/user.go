package request

import users "vanilla-florist/business/user"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEdit struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *UserLogin) ToUsecase() *users.User {
	return &users.User{
		Email:    user.Email,
		Password: user.Password,
	}
}

func (user *UserEdit) ToUsecase() *users.User {
	return &users.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
