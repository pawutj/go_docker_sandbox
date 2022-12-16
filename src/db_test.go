package go_docker_sandbox

import (
	"fmt"
	"testing"
)

func TestConnectDB(t *testing.T) {
	db := ConnectDB()

	if db == nil {
		t.Fatalf("ConnectDB failed")
	}

	rows, err := db.Query("SELECT * FROM product WHERE NAME = 'SomeProduct';")
	if err != nil {
		t.Fatalf("Select failed")
	}

	for rows.Next() {
		var id, price int
		var name string

		err = rows.Scan(&id, &name, &price)

		if err != nil {
			t.Fatalf("Error from scan")
		}

		fmt.Println(id, " ", name)

		if name != "SomeProduct" {
			t.Errorf("Want SomeProduct result %s", name)
		}
	}

}
