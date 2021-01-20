package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func BasicExample() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: selectColumn column <file1> [<file2> [... <fileN>]]\n")
		os.Exit(1)
	}

	temp, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("column value is not an integer:", temp)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("invalid column number")
		os.Exit(1)
	}

	for _, filename := range arguments[2:] {
		fmt.Println("\t\t", filename)

		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}

		defer f.Close()

		r := bufio.NewReader(f)

		for {
			line, err := r.ReadString('\n')
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}

				fmt.Printf("error reading file: %s\n", err)
			}

			data := strings.Fields(line)

			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}
}

func ChangeDT() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide one text file to process!")
		os.Exit(1)
	}

	filename := arguments[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}

	defer f.Close()

	notAMatch := 0

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			fmt.Printf("error reading file %s", err)
		}

		if r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`); r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)

			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err != nil {
				notAMatch++
				continue
			}

			newFormat := d1.Format(time.Stamp)
			fmt.Print(strings.Replace(line, match[1], newFormat, 1))

			continue
		}

		if r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`); r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)

			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err != nil {
				notAMatch++
				continue
			}

			newFormat := d1.Format(time.Stamp)
			fmt.Print(strings.Replace(line, match[1], newFormat, 1))

			continue
		}
	}

	fmt.Println(notAMatch, "lines did not match")
}

func FindIPv4() {

}

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"

	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP

	matchMe := regexp.MustCompile(grammar)

	return matchMe.FindString(input)
}

func main() {
	//BasicExample()
	//ChangeDT()
	FindIPv4()
}
