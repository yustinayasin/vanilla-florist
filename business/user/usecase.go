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

	userRepo, err := userUseCase.Repo.SignUp(user)

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

	userRepo, err := userUseCase.Repo.Login(user)

	if err != nil {
		return User{}, err
	}

	match := helpers.CheckPasswordHash(user.Password, userRepo.Password)

	if match != true {
		return User{}, errors.New("Password doesn't match")
	}

	userRepo.Token = userUseCase.Jwt.GenerateToken(userRepo.Id)

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

	userRepo, err := userUseCase.Repo.EditUser(user, id)

	if err != nil {
		return User{}, err
	}

	return userRepo, nil
}

func (userUseCase *UserUseCase) DeleteUser(id int) (User, error) {
	if id == 0 {
		return User{}, errors.New("User ID empty")
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

// func (userUseCase *UserUseCase) FindUser(id int) (User, error) {
// 	fmt.Println("Inside FindUser use case function")

// 	// Check if userUseCase or userUseCase.Repo is nil
// 	if userUseCase == nil {
// 		return User{}, errors.New("userUseCase is nil")
// 	}
// 	if userUseCase.Repo == nil {
// 		return User{}, errors.New("userUseCase.Repo is nil")
// 	}

// 	fmt.Println("Before calling FindUser on repository")

// 	// Call FindUser on the repository
// 	userRepo, err := userUseCase.Repo.FindUser(id)

// 	if err != nil {
// 		fmt.Println("Error calling FindUser on repository:", err)
// 		return User{}, err
// 	}

// 	fmt.Println("User found in repository:", userRepo)

// 	// If userRepo is nil, return an error indicating that the user was not found
// 	if userRepo == nil {
// 		fmt.Println("User not found in repository")
// 		return User{}, errors.New("user not found")
// 	}

// 	return userRepo, nil
// }
