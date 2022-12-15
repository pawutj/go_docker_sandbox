package go_docker_sandbox

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := ""
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Okay.")

}
