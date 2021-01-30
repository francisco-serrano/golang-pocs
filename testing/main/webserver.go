package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	body := "The current time is:"
	fmt.Fprintf(w, `<h1 align="center">%s</h1>`, body)

	now := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, `<h2 align="center">%s</h2>\n`, now)

	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func getData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)

	connStr := "user=postgres dbname=s2 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(w, `<h1 align="center">%s</h1>`, err)
		return
	}

	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Fprintf(w, `<h3 align="center">%s</h3>\n`, err)
		return
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		var id int
		var firstName, lastName string

		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			fmt.Fprintf(w, `<h1 align="center">%s</h1>\n`, err)
			return
		}

		fmt.Fprintf(w, `<h3 align="center"></h3>\n`, id, firstName, lastName)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(w, `<h1 align="center">%s</h1>`, err)
		return
	}
}

func main() {

}
