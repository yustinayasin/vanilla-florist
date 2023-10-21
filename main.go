package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	userUsecase "vanilla-florist/business/user"
	userController "vanilla-florist/controller/user"
	userRepo "vanilla-florist/drivers/databases/user"

	_ "github.com/lib/pq"
)

func main() {
	//connection
	connStr := "user=postgres dbname=florist password=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	//check if the connection work
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	userRepoInterface := userRepo.NewUserRepository(db)
	userUseCaseInterface := userUsecase.NewUseCase(userRepoInterface)
	userControllerInterface := userController.NewUserController(userUseCaseInterface)

	http.HandleFunc("/user/add", userControllerInterface.SignUp)
}
