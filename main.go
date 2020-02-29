package main

//go:generate sqlboiler --wipe psql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	_, err := sql.Open("postgres", `dbname=mari host=localhost user=postgres password=postgres`)
	dieIf(err)

	fmt.Println("connected")
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}