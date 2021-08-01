package main

import (
	"app/server"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	addr := flag.String("host", "localhost", "Address to listen on for the HTTP server")
	port := flag.Int("port", 8080, "Port for the HTTP server to listen on")
	version := flag.String("version", "", "Deployed version of the app")

	s := server.New(*version)
	srv := &http.Server{
		ReadTimeout:  2 * time.Second,   // Time to read the request
		WriteTimeout: 10 * time.Second,  // Time to write a response
		IdleTimeout:  120 * time.Second, // Max time for keep-alive waits
		Addr:         fmt.Sprintf("%s:%d", *addr, *port),
		Handler:      s,
	}

	log.Fatal(srv.ListenAndServe())
}
