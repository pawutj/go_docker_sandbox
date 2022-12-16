package go_docker_sandbox

import (
	"testing"
)

func TestConnectDB(t *testing.T) {
	db := ConnectDB()
	defer db.Exec("Drop table product")
	if db == nil {
		t.Fatalf("ConnectDB failed")
	}

	rows, err := db.Query("SELECT * FROM product WHERE NAME = 'SomeProduct';")
	if err != nil {
		t.Fatalf("Select failed")
	}

	var id, price int
	var name string

	rows.Next()
	err = rows.Scan(&id, &name, &price)

	if err != nil {
		t.Fatalf("Error from scan")
	}

	if name != "SomeProduct" {
		t.Errorf("Want SomeProduct result %s", name)
	}

}
