package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// se connecte à la base de données containérisée,
	db, err := sql.Open("mysql", "{user}:{password}@tcp(127.0.0.1:3306)/{dbname}")
	if err != nil {
		panic(err.Error())
	}
	// Close db :
	defer db.Close()

	// INSERT :
	insert, err := db.Query("INSERT INTO test VALUES ( 2, '' )")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
