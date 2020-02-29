package main

//go:generate sqlboiler --wipe psql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	_, _ := sql.Open("postgres", `dbname=mari host=localhost user=postgres password=postgres`)
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}