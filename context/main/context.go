package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1.A():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1.B():", r)
	}
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2.A():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2.B():", r)
	}
}

func f3(t int) {
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)

	c3 := context.Background()
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3.A():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3.B():", r)
	}
}

func exampleA() {
	delaySeconds := 6

	fmt.Println("Delay:", delaySeconds)

	f1(delaySeconds)
	f2(delaySeconds)
	f3(delaySeconds)
}

func exampleB() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://www.google.com:81", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

type aKey string

func searchKey(ctx context.Context, k aKey) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("found value:", v)
		return
	}

	fmt.Println("key not found:", k)
}

func exampleC() {
	myKey := aKey("mySecretValue")
	ctx := context.WithValue(context.Background(), myKey, "mySecretValue")

	searchKey(ctx, myKey)
	searchKey(ctx, "notThere")

	emptyCtx := context.TODO()

	searchKey(emptyCtx, "notThere")
}

func main() {
	//exampleA()
	//exampleB()
	exampleC()
}
