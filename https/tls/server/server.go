package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func main() {
	caCert, err := ioutil.ReadFile("client.crt")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	port := ":8080"

	srv := &http.Server{
		Addr:      port,
		Handler:   &handler{},
		TLSConfig: cfg,
	}

	fmt.Println("Listening on port number", port)
	fmt.Println(srv.ListenAndServeTLS("server.crt", "server.key"))
}
