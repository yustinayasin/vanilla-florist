package users

import (
	"errors"
	"vanilla-florist/helpers"
)

type UserUseCase struct {
	repo UserRepoInterface
}

func NewUseCase(userRepo UserRepoInterface) UserUseCaseInterface {
	return &UserUseCase{
		repo: userRepo,
	}
}

func (userUseCase *UserUseCase) SignUp(user User) (User, error) {
	if user.Name == "" {
		return User{}, errors.New("Name empty")
	}

	if user.Email == "" {
		return User{}, errors.New("Email empty")
	}

	if user.Password == "" {
		return User{}, errors.New("Password empty")
	}

	hash, _ := helpers.HashPassword(user.Password)

	user.Password = hash

	userRepo, err := userUseCase.repo.SignUp(user)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}
