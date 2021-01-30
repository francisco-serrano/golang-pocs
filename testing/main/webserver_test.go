package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createTable() {
	connStr := "user=postgres dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	const query = `create table if not exists users (
		id serial primary key,
		first_name text,
		last_name text
	)`

	if _, err := db.Exec(query); err != nil {
		fmt.Println(err)
		return
	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
	}
}

func dropTable() {
	connStr := "user=postgres dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err = db.Exec(`drop table if exists users`); err != nil {
		fmt.Println(err)
		return
	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
	}
}

func insertRecord(query string) {
	connStr := "user=postgres dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = db.Exec(query); err != nil {
		fmt.Println(err)
		return
	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
	}
}

func TestCount(t *testing.T) {
	createTable()
	insertRecord(`insert into users (first_name, last_name) values ('Epifanios', 'Doe'`)
	insertRecord(`insert into users (first_name, last_name) values ('Mihalis', 'Tsoukalos'`)
	insertRecord(`insert into users (first_name, last_name) values ('Mihalis', 'Unknown'`)

	connStr := "user=postgres dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	row := db.QueryRow(`select count(*) from users`)

	var count int
	if err := row.Scan(&count); err != nil {
		fmt.Println(err)
		return
	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
		return
	}

	if count != 3 {
		t.Errorf("Select query returned %d", count)
	}

	dropTable()
}

func TestQueryDB(t *testing.T) {
	createTable()

	connStr := "user=postgres dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	insertRecord(`insert into users (first_name, last_name) values ('Random Text', '123456')`)

	rows, err := db.Query(`select * from users where last_name=$1`, "123456")
	if err != nil {
		fmt.Println(err)
		return
	}

	var col1 int
	var col2, col3 string

	for rows.Next() {
		if err := rows.Scan(&col1, &col2, &col3); err != nil {
			fmt.Println(err)
		}
	}

	if col2 != "Random Text" {
		t.Errorf("first_name returned %s", col2)
	}

	if col3 != "123456" {
		t.Errorf("last_name returned %s", col3)
	}

	if err := db.Close(); err != nil {
		fmt.Println(err)
	}

	dropTable()
}

func TestRecord(t *testing.T) {
	createTable()
	insertRecord(`insert into users (first_name, last_name) values ('John', 'Doe')`)

	req, err := http.NewRequest(http.MethodGet, "/getdata", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(getData)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned %v", status)
	}

	if rr.Body.String() != `<h3 align="center">1, John, Doe</h3>\n` {
		t.Errorf("Wrong server response!")
	}

	dropTable()
}