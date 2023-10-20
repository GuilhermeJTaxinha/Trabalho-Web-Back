package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectDBPostgres() (*sql.DB, error) {
	var err error

	//dsn := fmt.Sprintf("host=localhost port=5432 user=postgres password=minhasenha123 dbname=projeto sslmode=disable")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=minhasenha123 dbname=projeto sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect DB")
	}

	err = db.Ping()

	return db, err
}

//postgres", "host=localhost port=5432 user=postgres password=minhasenha123 dbname=projeto sslmode=disable
