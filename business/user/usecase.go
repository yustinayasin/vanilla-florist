package users

import (
	"errors"
	"vanilla-florist/helpers"
)

type GeneratorToken interface {
	GenerateToken(userId int) string
}

type UserUseCase struct {
	Repo UserRepoInterface
	Jwt  GeneratorToken
}

func NewUseCase(userRepo UserRepoInterface, tokenGenerator GeneratorToken) UserUseCaseInterface {
	return &UserUseCase{
		Repo: userRepo,
		Jwt:  tokenGenerator,
	}
}

func (userUseCase *UserUseCase) SignUp(user User) (User, error) {
	if user.Name == "" {
		return User{}, errors.New("name cannot be empty")
	}

	if user.Email == "" {
		return User{}, errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return User{}, errors.New("password cannot be empty")
	}

	hash, _ := helpers.HashPassword(user.Password)

	user.Password = hash

	userRepo, err := userUseCase.Repo.SignUp(user)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) Login(user User) (User, error) {
	if user.Email == "" {
		return User{}, errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return User{}, errors.New("password cannot be empty")
	}

	userRepo, err := userUseCase.Repo.Login(user)

	if err != nil {
		return User{}, err
	}

	match := helpers.CheckPasswordHash(user.Password, userRepo.Password)

	if match != true {
		return User{}, errors.New("password doesn't match")
	}

	userRepo.Token = userUseCase.Jwt.GenerateToken(userRepo.Id)

	return userRepo, nil
}

func (userUseCase *UserUseCase) EditUser(user User, id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("user ID cannot be empty")
	}

	if user.Name == "" {
		return User{}, errors.New("name cannot be empty")
	}

	if user.Email == "" {
		return User{}, errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return User{}, errors.New("password cannot be empty")
	}

	hash, _ := helpers.HashPassword(user.Password)
	user.Password = hash

	userRepo, err := userUseCase.Repo.EditUser(user, id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) DeleteUser(id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("user ID cannot be empty")
	}

	userRepo, err := userUseCase.Repo.DeleteUser(id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) FindUser(id int) (User, error) {
	userRepo, err := userUseCase.Repo.FindUser(id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}
