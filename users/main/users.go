package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	fmt.Println("user id:", os.Getuid())

	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Group ids: ")

	groupIDs, err := u.GroupIds()
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range groupIDs {
		fmt.Print(i, " ")
	}

	fmt.Println()
}
