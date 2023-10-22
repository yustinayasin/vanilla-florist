package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vanilla-florist/helpers"

	userUsecase "vanilla-florist/business/user"
	userController "vanilla-florist/controller/user"
	userRepo "vanilla-florist/drivers/databases/user"

	_ "github.com/lib/pq"
)

func main() {
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
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	http.HandleFunc("/user/add", userControllerInterface.SignUp)

	// listen port
	err = http.ListenAndServe(":3000", nil)

	// print any server-based error messages
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
