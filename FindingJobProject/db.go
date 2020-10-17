package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func connectDatabse() {

	db, err = sql.Open("mysql", "root:Quang123Huy@@/vieclam")
	fmt.Println("Database connected.")
	if err != nil {
		panic(err.Error())
	}
}
