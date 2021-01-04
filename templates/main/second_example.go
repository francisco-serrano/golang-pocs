package main

import (
	"html/template"
	"log"
	"net/http"
)

var testTemplate *template.Template

type Widget struct {
	Name  string
	Price int
}

type ViewData struct {
	Name    string
	Widgets []Widget
}

func RunSecondExample() {
	var err error
	testTemplate, err = template.ParseFiles("./templates/hello.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := ViewData{
		Name:    "John Smith",
		Widgets: []Widget{
			{"Blue Widget", 12},
			{"Red Widget", 12},
			{"Green Widget", 12},
		},
	}

	if err := testTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
