package main

import (
	"app/server"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Host the HTTP server should bind to
var Host string = "localhost"

// Port the HTTP server should bind to
var Port string = "8080"

// Version of the application exposed on /version route
var Version string = "0.0.0"

func main() {

	s := server.New(Version)
	srv := &http.Server{
		ReadTimeout:  2 * time.Second,   // Time to read the request
		WriteTimeout: 10 * time.Second,  // Time to write a response
		IdleTimeout:  120 * time.Second, // Max time for keep-alive waits
		Addr:         fmt.Sprintf("%s:%s", Host, Port),
		Handler:      s,
	}

	log.Fatal(srv.ListenAndServe())
}
