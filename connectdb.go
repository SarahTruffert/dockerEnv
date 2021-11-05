// Vous devrez aussi fournir un petit programme
// (dans le langage de programmation que vous utilisez en entreprise)
// qui va se connecter à votre base de données containérisée,
// et y effectuer quelques opérations : au minimum "SELECT" et "INSERT".
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	db, err := sql.Open("mysql", "sarah:123Codons@tcp(127.0.0.1:3306)/sarahbdd")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}
