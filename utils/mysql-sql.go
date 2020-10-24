package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "bookstore:Bookstore.com@_00..@tcp(62.234.11.179:3306)/bookstore")
	if err != nil {
		panic(err.Error())
	}
}
