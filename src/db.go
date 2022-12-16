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

	// createProductTable := `CREATE TABLE IF NOT EXISTS product (id SERIAL  PRIMARY KEY , name TEXT , price INTEGER);`
	// _, err = db.Exec(createProductTable)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// row := db.QueryRow("Insert into product (name, price) values ($1 , $2) RETURNING ID", "SomeProduct", 1)
	// var id int
	// err = row.Scan(&id)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("insert complate", id)

	return db
}
