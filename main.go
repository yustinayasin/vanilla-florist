package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vanilla-florist/app/middleware"
	"vanilla-florist/helpers"

	userUsecase "vanilla-florist/business/user"
	userController "vanilla-florist/controller/user"
	userRepo "vanilla-florist/drivers/databases/user"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration from config.json
	config, err := middleware.LoadConfig("config.json")

	if err != nil {
		log.Fatal(err)
	}

	db, err := helpers.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	//check if the connection work
	pingErr := db.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	userRepoInterface := userRepo.NewUserRepository(db)
	jwtConf := middleware.ConfigJWT{
		SecretJWT:       config.SecretJWT,
		ExpiresDuration: config.ExpiresDuration,
	}
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface, jwtConf)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	r := mux.NewRouter()

	r.HandleFunc("/user/signup", userControllerInterface.SignUp)
	r.HandleFunc("/user/login", userControllerInterface.Login)

	// Protected routes using RequireAuth middleware
	r.HandleFunc("/user/edit/{id}", middleware.RequireAuth(userControllerInterface.EditUser, jwtConf, userRepoInterface))
	r.HandleFunc("/user/delete/{id}", middleware.RequireAuth(userControllerInterface.DeleteUser, jwtConf, userRepoInterface))

	// listen port
	err = http.ListenAndServe(":3000", r)

	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
