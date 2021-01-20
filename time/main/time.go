package main

import (
	"fmt"
	"github.com/thanhpk/randstr"
	"log"
	"regexp"
	"time"
)

func SampleA() {
	a := "20 July 2000"

	d, err := time.Parse("02 January 2006", a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)
	fmt.Println(d)

	logs := []string{"127.0.0.1 - - [16/Nov/2017:10:49:46 +0200]  325504", "127.0.0.1 - - [16/Nov/2017:10:16:41 +0200] \"GET /CVEN  HTTP/1.1\" 200 12531 \"-\" \"Mozilla/5.0 AppleWebKit/537.36", "127.0.0.1 200 9412 - - [12/Nov/2017:06:26:05 +0200]  \"GET \"http://www.mtsoukalos.eu/taxonomy/term/47\" 1507",
		"[12/Nov/2017:16:27:21 +0300]",
		"[12/Nov/2017:20:88:21 +0200]",
		"[12/Nov/2017:20:21 +0200]",
	}

	for _, entry := range logs {
		r := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\].*`)

		if r.MatchString(entry) {
			match := r.FindStringSubmatch(entry)

			dt, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err != nil {
				fmt.Println("not a valid date time format")
			} else {
				newFormat := dt.Format(time.RFC850)
				fmt.Println(newFormat)
			}
		} else {
			fmt.Println("not a match!")
		}
	}
}

func main() {
	//SampleA()
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("20060102"))
	fmt.Println(randstr.String(10))
}
