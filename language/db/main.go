package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func main () {
	strConnection := "root:docker@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!", db)

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully pinged the database!")

	// Example query
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println(rows)
}
