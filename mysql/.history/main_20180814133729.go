package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/research")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("SELECT * from test")
	if err != nil {
		panic(err.Error())
	}
}