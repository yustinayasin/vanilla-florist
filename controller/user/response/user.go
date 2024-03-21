package response

import (
	users "vanilla-florist/business/user"
)

type UserResponse struct {
	Id			int    `json:"id"`
	Email		string `json:"email"`
	Name		string `json:"name"`
	Password	string `json:"password"`
	Token		string `json:"token"`
}

func FromUsecase(user users.User) UserResponse {
	return UserResponse{
		Id:			user.Id,
		Name:		user.Name,
		Email:		user.Email,
		Password:	user.Password,
		Token:		user.Token
	}
}

func FromUsecaseList(user []users.User) []UserResponse {
	var userResponse []UserResponse

	for _, v := range user {
		userResponse = append(userResponse, FromUsecase(v))
	}

	return userResponse
}
