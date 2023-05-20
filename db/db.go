package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func DataBaseConection() *sql.DB {
	conection := "user=postgres dbname=bazar_loja password=220800 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conection)
	if err != nil {
		panic(err)
	}
	return db
}
