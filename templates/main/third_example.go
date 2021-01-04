package main

import (
	"html/template"
	"log"
	"net/http"
)

var testTemplate2 *template.Template

type User struct {
	Admin bool
}

type ViewData2 struct {
	*User
}

func RunThirdExample() {
	var err error
	testTemplate2, err = template.ParseFiles("./templates/hello2.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler2)
	http.ListenAndServe(":3000", nil)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := ViewData2{&User{true}}

	if err := testTemplate2.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
