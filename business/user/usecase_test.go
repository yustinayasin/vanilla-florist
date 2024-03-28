package users_test

import (
	"errors"
	"reflect"
	"testing"
	users "vanilla-florist/business/user"
	userRepo "vanilla-florist/drivers/databases/user"
)

var userUseCaseInterface users.UserUseCaseInterface
var userDataDummyLogin, userDataDummyEdit users.User

// Custom UserRepository implementation for testing
type MockUserRepository struct{}

func (r *MockUserRepository) Login(user users.User) (users.User, error) {
	// Simulate an error in the database
	return users.User{}, errors.New("error in database")
}

func (r *MockUserRepository) SignUp(users.User) (users.User, error) {
	// Simulate an error in the database
	return users.User{}, errors.New("error in database")
}

func (r *MockUserRepository) EditUser(users.User) (users.User, error) {
	// Simulate an error in the database
	return users.User{}, errors.New("error in database")
}

func (r *MockUserRepository) DeleteUser(id int) (users.User, error) {
	// Simulate successful deletion of user with given ID
	return users.User{
		Id: id,
		// Populate other fields as needed
	}, nil
}

func setup() {
	//data mock hasil login
	userDataDummyLogin = users.User{
		Id:       3,
		Name:     "Jeong Jeahyun",
		Email:    "jeongjaehyun@gmail.com",
		Password: "1234",
		Token:    "",
	}

	userDataDummyEdit = users.User{
		Id:       4,
		Name:     "Lee Mark",
		Email:    "leemark@gmail.com",
		Password: "1234",
	}
}

func AssertEqual(t *testing.T, expected, actual interface{}) {
	// If both expected and actual are errors, compare their string representations
	if expectedErr, ok := expected.(error); ok {
		if actualErr, ok := actual.(error); ok {
			if expectedErr.Error() != actualErr.Error() {
				t.Errorf("expected error: %q, got: %q", expectedErr, actualErr)
			}
			return
		}
	}

	// For other types, use default comparison
	if expected != actual {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}
}

func AssertError(t *testing.T, err error) bool {
	if err == nil {
		t.Helper()
		t.Errorf("Expected an error, but got nil")
		return false
	}
	return true
}

func TestLogin(t *testing.T) {
	setup()

	// subtest
	t.Run("Email empty", func(t *testing.T) {
		requestLoginUser := users.User{
			Email:    "",
			Password: "123",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.Login(requestLoginUser)

		AssertEqual(t, errors.New("email cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})

	t.Run("Password empty", func(t *testing.T) {
		requestLoginUser := users.User{
			Email:    "jeongjaehyun@gmail.com",
			Password: "",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.Login(requestLoginUser)

		AssertEqual(t, errors.New("password cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})

	t.Run("Password doesn't match", func(t *testing.T) {
		// Create a new instance of UserUseCase with the real repository
		useCase := users.NewUseCase(userRepo.NewUserRepository(nil), nil)

		// Call the login function
		user, err := useCase.Login(users.User{
			Email:    "jeongjaehyun@gmail.com",
			Password: "1234",
		})

		// Assert that an error occurred during login
		AssertError(t, err)

		// Assert that the error message matches
		expectedErrorMsg := "password doesn't match"
		if err != nil && err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: %q, got: %q", expectedErrorMsg, err.Error())
			return
		}

		// Ensure that the user object is empty
		AssertEqual(t, users.User{}, user)
	})

	t.Run("User not found", func(t *testing.T) {
		// Create a new instance of UserUseCase with the real repository
		useCase := users.NewUseCase(userRepo.NewUserRepository(nil), nil)

		// Call the login function
		user, err := useCase.Login(users.User{
			Email:    "jeongjaehyunn@gmail.com",
			Password: "1234",
		})

		// Assert that an error occurred during login
		AssertError(t, err)

		// Assert that the error message matches
		expectedErrorMsg := "User not found"
		if err != nil && err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: %q, got: %q", expectedErrorMsg, err.Error())
			return
		}

		// Ensure that the user object is empty
		AssertEqual(t, users.User{}, user)
	})
}

func TestSignUp(t *testing.T) {
	setup()

	// subtest
	t.Run("Email empty", func(t *testing.T) {
		requestSignUpUser := users.User{
			Name:     "Jeong Jaehyun",
			Email:    "",
			Password: "123",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.SignUp(requestSignUpUser)

		AssertEqual(t, errors.New("email cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})

	t.Run("Password empty", func(t *testing.T) {
		requestSignUpUser := users.User{
			Name:     "Jeong Jaehyun",
			Email:    "jeongjaehyun@gmail.com",
			Password: "",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.SignUp(requestSignUpUser)

		AssertEqual(t, errors.New("password cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})

	t.Run("Name empty", func(t *testing.T) {
		requestSignUpUser := users.User{
			Name:     "",
			Email:    "jeongjaehyun@gmail.com",
			Password: "123",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.SignUp(requestSignUpUser)

		AssertEqual(t, errors.New("name cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})
}

func TestEdit(t *testing.T) {
	setup()

	t.Run("Edit user with empty ID", func(t *testing.T) {
		// Create a new instance of UserUseCase with the real repository
		useCase := users.NewUseCase(userRepo.NewUserRepository(nil), nil)

		// Define a dummy user with empty ID
		dummyUser := users.User{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Password: "password",
		}

		// Call the EditUser function with empty ID
		editedUser, err := useCase.EditUser(dummyUser, 0)

		// Assert that an error occurred due to empty ID
		AssertError(t, err)

		// Assert that the error message matches
		expectedErrorMsg := "user ID cannot be empty"
		if err != nil && err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: %q, got: %q", expectedErrorMsg, err.Error())
			return
		}

		// Assert that the editedUser is empty
		AssertEqual(t, users.User{}, editedUser)
	})

	t.Run("Email empty", func(t *testing.T) {
		dummyUser := users.User{
			Name:     "Jeong Jaehyun",
			Email:    "",
			Password: "123",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.EditUser(dummyUser, 4)

		AssertEqual(t, errors.New("email cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})

	t.Run("Password empty", func(t *testing.T) {
		dummyUser := users.User{
			Name:     "Jeong Jaehyun",
			Email:    "jeongjaehyun@gmail.com",
			Password: "",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.EditUser(dummyUser, 4)

		AssertEqual(t, errors.New("password cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})

	t.Run("Name empty", func(t *testing.T) {
		dummyUser := users.User{
			Name:     "",
			Email:    "jeongjaehyun@gmail.com",
			Password: "123",
		}

		userUseCaseInterface = &users.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.EditUser(dummyUser, 4)

		AssertEqual(t, errors.New("name cannot be empty"), err)
		AssertEqual(t, users.User{}, user)
	})
}

func TestDelete(t *testing.T) {
	setup()

	t.Run("Delete user with empty ID", func(t *testing.T) {
		// Create a new instance of UserUseCase with the real repository
		useCase := users.NewUseCase(userRepo.NewUserRepository(nil), nil)

		// Call the EditUser function with empty ID
		editedUser, err := useCase.DeleteUser(0)

		// Assert that an error occurred due to empty ID
		AssertError(t, err)

		// Assert that the error message matches
		expectedErrorMsg := "user ID cannot be empty"
		if err != nil && err.Error() != expectedErrorMsg {
			t.Errorf("Expected error message: %q, got: %q", expectedErrorMsg, err.Error())
			return
		}

		// Assert that the editedUser is empty
		AssertEqual(t, users.User{}, editedUser)
	})

	t.Run("Success delete", func(t *testing.T) {
		// Create a new instance of UserUseCase with the real repository
		useCase := users.NewUseCase(userRepo.NewUserRepository(nil), nil)

		// Call the EditUser function with empty ID
		deletedUser, err := useCase.DeleteUser(4)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			return
		}

		// Assert that the deletedUser matches the expected user
		expectedDeletedUser := users.User{
			// Populate the expected fields based on your business logic
			// For example:
			Id:       3,               // Assuming the deleted user's ID is 3
			Name:     "Jeong Jaehyun", // Assuming the name of the deleted user
			Email:    "jeongjaehyun@gmail.com",
			Password: "password123",
		}

		if !reflect.DeepEqual(deletedUser, expectedDeletedUser) {
			t.Errorf("Expected deleted user: %+v, but got: %+v", expectedDeletedUser, deletedUser)
			return
		}
	})
}
