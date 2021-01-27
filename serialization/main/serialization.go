package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
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

func main() {
	if err := load(); err != nil {
		fmt.Println(err)
	}

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
