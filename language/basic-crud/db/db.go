package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes a connection to the MySQL database and returns the connection object.
func Connect() (*sql.DB, error) {
	strConnection := "root:docker@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}