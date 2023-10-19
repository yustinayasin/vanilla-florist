package main

import (
	"database/sql"
    "fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//connection
	connStr := "user=postgres dbname=florist password=Mudjinah19 sslmode=disable"

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
}

