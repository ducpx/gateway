package main

import (
	"log"
	"net/http"

	"github.com/ducpx/proxy"
	"github.com/nats-io/nats.go"
)

func main() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	px, err := proxy.NewNatsProxy(conn)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.Handle("/", px)

	http.ListenAndServe(":3000", mux)
}
