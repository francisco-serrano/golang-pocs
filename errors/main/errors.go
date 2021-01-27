package main

import (
	"fmt"
	"log"
)

func getEndpoint() error {
	if err := getElementFromRepository(); err != nil {
		return fmt.Errorf("error while getting from repository: %w", err)
	}

	return nil
}

func getElementFromRepository() error {
	if err := getElementFromDB(); err != nil {
		return fmt.Errorf("error while getting from DB: %w", err)
	}

	return nil
}

func getElementFromDB() error {
	return fmt.Errorf("record not found")
}

func main() {
	if err := getEndpoint(); err != nil {
		log.Fatal(err)
	}
}
