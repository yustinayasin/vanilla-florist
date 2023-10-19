package user

import (
	"context"
	"errors"
	"time"
	"net/http"
)

func SignUp(user User) (User, error) {
	if user.Email == "" {
		return User{}, errors.New("Email empty")
	}

	if user.Password == "" {
		return User{}, errors.New("Password empty")
	}

	hash, _ := helpers.HashPassword(user.Password)

	user.Password = hash

	userRepo, err := userUseCase.repo.SignUp(user, ctx)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}