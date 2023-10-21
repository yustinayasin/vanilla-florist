package controllers

import (
	"encoding/json"
	"net/http"

	users "vanilla-florist/business/user"
	"vanilla-florist/controller/user/request"
	"vanilla-florist/utils"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}

// dipasangkan dengan routing
func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}
}

// Add a user handler
func (controller *UserController) SignUp(res http.ResponseWriter, req *http.Request) {
	// check the method
	if req.Method != "POST" {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
	}

	var userSignup request.UserLogin

	payload := req.Body

	//defer ensure req.Body.Close() will be executed after the AddMovie or schedule a function
	defer req.Body.Close()

	// parse the movie data into json format
	err := json.NewDecoder(payload).Decode(&userSignup)

	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error parsing the movie data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	}

	_, errRepo := controller.usecase.SignUp(*userSignup.ToUsecase())

	if errRepo != nil {
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error in repo,
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	}

	// Add the response return message
	HandlerMessage := []byte(`{
	 "success": true,
	 "message": "User was successfully created",
	 }`)

	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
}
