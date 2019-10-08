package models

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:laravel@tcp(localhost:3306)/laravel")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	con = db
	return db
}
