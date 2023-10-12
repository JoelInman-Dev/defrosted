package data

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
var DB *sql.DB

func Connect() {
	dsn := "host=172.23.0.2 port=5432 user=defrosted password=defrosted dbname=defrosted sslmode=disable"
	var err error
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	
	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		fmt.Println("Error: Could not establish a connection with the database", err)
	}
}