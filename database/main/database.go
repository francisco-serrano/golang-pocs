package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionUrl := "root:root@(localhost:3306)/rely?charset=utf8&parseTime=True"

	db, err := sql.Open("mysql", connectionUrl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Ping())
	fmt.Printf("%+v\n", db.Stats())

	rows, err := db.Query("show tables")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rows.Columns())
}
