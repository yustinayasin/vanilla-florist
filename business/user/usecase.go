package users

import (
	"errors"
	"fmt"
	"vanilla-florist/app/middleware"
	"vanilla-florist/helpers"
)

type UserUseCase struct {
	repo UserRepoInterface
	jwt  middleware.ConfigJWT
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

func (userUseCase *UserUseCase) Login(user User) (User, error) {
	if user.Email == "" {
		return User{}, errors.New("Email cannot be empty")
	}

	if user.Password == "" {
		return User{}, errors.New("Password cannot be empty")
	}

	userRepo, err := userUseCase.repo.Login(user)

	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	match := helpers.CheckPasswordHash(user.Password, userRepo.Password)

	if match != true {
		return User{}, errors.New("Password doesn't match")
	}

	userRepo.Token = userUseCase.jwt.GenerateToken(user.Id)

	return userRepo, nil
}

func (userUseCase *UserUseCase) EditUser(user User, id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("User ID empty")
	}

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

	userRepo, err := userUseCase.repo.EditUser(user, id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) DeleteUser(id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("User ID empty")
	}

	userRepo, err := userUseCase.repo.DeleteUser(id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}
