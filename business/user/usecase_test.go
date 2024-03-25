package users_test

import (
	"errors"
	"testing"
	userUsecase "vanilla-florist/business/user"
)

var userUseCaseInterface userUsecase.UserUseCaseInterface
var userDataDummyLogin, userDataDummyEdit userUsecase.User

func setup() {
	//data mock hasil login
	userDataDummyLogin = userUsecase.User{
		Id:       3,
		Name:     "Lee Mark",
		Email:    "leemark@gmail.com",
		Password: "1234",
		Token:    "",
	}

	userDataDummyEdit = userUsecase.User{
		Id:       3,
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

func TestLogin(t *testing.T) {
	setup()

	// subtest
	t.Run("Email empty", func(t *testing.T) {
		requestLoginUser := userUsecase.User{
			Email:    "",
			Password: "123",
		}

		userUseCaseInterface = &userUsecase.UserUseCase{}

		// pake interface karena beda package
		user, err := userUseCaseInterface.Login(requestLoginUser)

		// if err != nil {
		// 	t.Fatalf("Email is empty!")
		// }

		AssertEqual(t, errors.New("Email cannot be empty"), err)
		AssertEqual(t, userUsecase.User{}, user)
	})
}
