package main

import (
	"fmt"
	"log"
)

func GetEndpoint() error {
	if err := GetElementFromRepository(); err != nil {
		return fmt.Errorf("error while getting from repository: %w", err)
	}

	return nil
}

func GetElementFromRepository() error {
	if err := GetElementFromDB(); err != nil {
		return fmt.Errorf("error while getting from DB: %w", err)
	}

	return nil
}

func GetElementFromDB() error {
	return fmt.Errorf("record not found")
}

func main() {
	if err := GetEndpoint(); err != nil {
		log.Fatal(err)
	}
}
