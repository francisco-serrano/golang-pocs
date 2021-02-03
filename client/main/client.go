package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func main() {
	url := "https://www.google.com"

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	httpData, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Status Code:", httpData.Status)

	header, err := httputil.DumpResponse(httpData, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")

	charset := strings.SplitAfter(contentType, "charset=")
	if len(charset) > 1 {
		fmt.Println("Charset:", charset[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}


	data, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer data.Body.Close()

	if _, err := io.Copy(os.Stdout, data.Body); err != nil {
		fmt.Println(err)
		return
	}
}
