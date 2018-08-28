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

	for rows.Next() {
		var id int
		var name string
		var address string
		var tel string
		err = rows.Scan(&id, &name, &address, &tel)
		if err != nil {
			panic(err.Error())
		}
	}
}