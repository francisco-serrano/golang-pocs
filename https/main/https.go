package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func exampleA() {
	url := "https://www.google.com"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{},
	}

	client := &http.Client{
		Transport: tr,
	}

	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	s := strings.TrimSpace(string(content))

	fmt.Println(s)
}

func exampleB() {
	port := ":8080"

	http.HandleFunc("/", Default)

	if err := http.ListenAndServeTLS(port, "server.crt", "server.key", nil); err != nil {
		log.Fatal(err)
	}
}

func Default(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is an example HTTPS server!\n")
}

func main() {
	//exampleA()
	exampleB()
}