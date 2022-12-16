package go_docker_sandbox

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	url := "postgres://root:root@db/go-example-db?sslmode=disable"
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
