package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var (
	DATA     = make(map[string]myElement)
	DATAFILE = "/tmp/dataFile.gob"
)

type myElement struct {
	Name    string
	Surname string
	ID      string
}

func save() error {
	fmt.Println("SAVING", DATAFILE)

	if err := os.Remove(DATAFILE); err != nil {
		fmt.Println(err)
	}

	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("cannot create", DATAFILE)
		return err
	}

	defer func() {
		if err := saveTo.Close(); err != nil {
			fmt.Println("cannot close", DATAFILE)
			fmt.Println(err)
		}
	}()

	encoder := gob.NewEncoder(saveTo)

	if err := encoder.Encode(DATA); err != nil {
		fmt.Println("cannot save to", DATAFILE)
		return err
	}

	return nil
}

func load() error {
	fmt.Println("loading", DATAFILE)

	loadFrom, err := os.Open(DATAFILE)
	if err != nil {
		fmt.Println("empty key/value store!")
		return err
	}

	defer func() {
		if err := loadFrom.Close(); err != nil {
			fmt.Println("cannot close", DATAFILE)
			fmt.Println(err)
		}
	}()

	decoder := gob.NewDecoder(loadFrom)
	if err := decoder.Decode(&DATA); err != nil {
		fmt.Println("cannot load from", DATAFILE)
		return err
	}

	return nil
}

func Add(k string, n myElement) bool {
	if k == "" {
		return false
	}

	if Lookup(k) == nil {
		DATA[k] = n
		return true
	}

	return false
}

func Delete(k string) bool {
	if Lookup(k) != nil {
		delete(DATA, k)
		return true
	}

	return false
}

func Lookup(k string) *myElement {
	n, ok := DATA[k]
	if !ok {
		return nil
	}

	return &n
}

func Change(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func Print() {
	for k, d := range DATA {
		fmt.Printf("key: %s value: %v\n", k, d)
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving", r.Host, "for", r.URL.Path)

	t := template.Must(template.ParseGlob("home.gohtml"))
	t.ExecuteTemplate(w, "home.gohtml", nil)
}

func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store!")

	fmt.Fprintf(w, `<a href="/" style="margin-right: 20px;">Home sweet home!</a>`)
	fmt.Fprintf(w, `<a href="/list" style="margin-right: 20px;">List all elements!</a>`)
	fmt.Fprintf(w, `<a href="/change" style="margin-right: 20px;">Change an element!</a>`)
	fmt.Fprintf(w, `<a href="/insert" style="margin-right: 20px;">Insert new element!</a>`)

	fmt.Fprintf(w, `<h1>The contents of the KV store are:</h1>`)

	fmt.Fprintf(w, `<ul>`)

	for k, v := range DATA {
		fmt.Fprintf(w, `<li>`)
		fmt.Fprintf(w, `<strong>%s</strong> with value: %v`, k, v)
		fmt.Fprintf(w, `</li>`)
	}

	fmt.Fprintf(w, `</ul>`)
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of the KV store!")

	t := template.Must(template.ParseFiles("update.gohtml"))
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}

	if !Change(key, n) {
		fmt.Println("Update operation failed!")
	} else {
		if err := save(); err != nil {
			fmt.Println(err)
			return
		}

		t.Execute(w, struct {
			Success bool
		}{true})
	}
}

func insertElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the KV store!")
	t := template.Must(template.ParseFiles("insert.gohtml"))
	if r.Method != http.MethodPost {
		t.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	n := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}

	if !Add(key, n) {
		fmt.Println("Add operation failed!")
		return
	}

	if err := save(); err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, struct {
		Success bool
	}{true})
}

func cliOperation() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)

		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			Print()
		case "STOP":
			if err := save(); err != nil {
				fmt.Println(err)
			}
			return
		case "DELETE":
			if !Delete(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !Add(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			if n := Lookup(tokens[1]); n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "CHANGE":
			n := myElement{tokens[2], tokens[3], tokens[4]}
			if !Change(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again!")
		}
	}
}

func webserverOperation() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func main() {
	if err := load(); err != nil {
		fmt.Println(err)
	}

	//cliOperation()
	webserverOperation()
}
