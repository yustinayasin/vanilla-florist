package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
		return
	}

	var userSignup request.UserEdit

	// Read the request body
	bodyBytes, err := ioutil.ReadAll(req.Body)

	if err != nil {
		// Handle error
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error read request body",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	// Create a new io.ReadCloser with the same content
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// parse the user data into json format
	err = json.NewDecoder(req.Body).Decode(&userSignup)

	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error parsing the user data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	//defer ensure req.Body.Close() will be executed after the AddMovie or schedule a function
	defer req.Body.Close()

	_, errRepo := controller.usecase.SignUp(*userSignup.ToUsecase())

	if errRepo != nil {
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error in repo,
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	// Add the response return message
	HandlerMessage := []byte(`{
	 "success": true,
	 "message": "User was successfully created",
	 }`)

	utils.ReturnJsonResponse(res, http.StatusCreated, HandlerMessage)
	return
}

func (controller *UserController) Login(res http.ResponseWriter, req *http.Request) {
	// check the method
	if req.Method != "GET" {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Check your HTTP method: Invalid HTTP method executed",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	var userLogin request.UserLogin

	// Read the request body
	bodyBytes, err := ioutil.ReadAll(req.Body)

	if err != nil {
		// Handle error
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error read request body",
		}`)

		utils.ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}

	// Create a new io.ReadCloser with the same content
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// parse the user data into json format
	err = json.NewDecoder(req.Body).Decode(&userLogin)

	if err != nil {
		// Add the response return message
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error parsing the user data",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	//defer ensure req.Body.Close() will be executed after the AddMovie or schedule a function
	defer req.Body.Close()

	_, errRepo := controller.usecase.Login(*userLogin.ToUsecase())

	if errRepo != nil {
		HandlerMessage := []byte(`{
		"success": false,
		"message": "Error query the user",
		}`)

		utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
		return
	}

	HandlerMessage := []byte(`{
		"success": success,
		"message": "Login success!",
	}`)

	utils.ReturnJsonResponse(res, http.StatusInternalServerError, HandlerMessage)
	return
}
