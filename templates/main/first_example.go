package main

import (
	"html/template"
	"log"
	"os"
)

type Test struct {
	HTML     string
	SafeHTML template.HTML
	Title    string
	Path     string
	Dog      Dog
	Map      map[string]string
}

type Dog struct {
	Name string
	Age int
}

func Run() {
	t, err := template.ParseFiles("./templates/context.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	data := Test{
		HTML:     "<h1>A header!</h1>",
		SafeHTML: template.HTML("<h1>A Safe header</h1>"),
		Title:    "Backslash! An in depth look at the \"\\\" character.",
		Path:     "/dashboard/settings",
		Dog:      Dog{
			Name: "Fido",
			Age:  6,
		},
		Map: map[string]string{
			"key": "value",
			"other_key": "other_value",
		},
	}

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
