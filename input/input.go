package input

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	scanner()
	//cliArguments()
}

func scanner() {
	f := os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}