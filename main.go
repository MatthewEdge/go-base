package main

import (
	"app/server"
	"net/http"
)

func main() {
	s := server.New()
	http.ListenAndServe("0.0.0.0:8080", s)
}
