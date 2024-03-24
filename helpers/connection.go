package helpers

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	connStr := "user=admin dbname=florist password=password sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
