package files

import (
	"fmt"
	"os"
)

func Run() {
	dirName := "./files"

	f, err := os.Open(dirName)
	if err != nil {
		panic(err)
	}

	files, err := f.Readdir(-1)
	if err != nil {
		panic(err)
	}

	if err = f.Close(); err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
