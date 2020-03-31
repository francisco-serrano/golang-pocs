package main

import (
	"flag"
	"fmt"
	"github.com/francisco-serrano/awesomeProject/context/log"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")

	fmt.Printf("value for foo is %v", ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
