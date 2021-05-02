package main

import (
	"app/server"
	"log"
	"net/http"
	"time"
)

func main() {
	s := server.New()
	srv := &http.Server{
		ReadTimeout:  2 * time.Second,   // Time to read the request
		WriteTimeout: 10 * time.Second,  // Time to write a response
		IdleTimeout:  120 * time.Second, // Max time for keep-alive waits
		Handler:      s,
		Addr:         "0.0.0.0:8080",
		// TLSConfig: &tls.Config{
		// // Use Go's default ciphersuite preferences to avoid most attacks
		// PreferServerCipherSuites: true,
		// // Mozilla Modern recommended suites. https://wiki.mozilla.org/Security/Server_Side_TLS
		// MinVersion: tls.VersionTLS13,
		// CipherSuites: []uint16{
		// tls.TLS_AES_128_GCM_SHA256,
		// tls.TLS_AES_256_GCM_SHA384,
		// tls.TLS_CHACHA20_POLY1305_SHA256,
		// },
		// },
	}

	log.Fatal(srv.ListenAndServe())
}
