package main

import (
	"app/env"
	"app/server"
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Host the HTTP server should bind to
var Host string = "localhost"

// Port the HTTP server should bind to
var Port string = "8080"

// Version of the application exposed on /version route
var Version string = "0.0.0"

//go:embed migrations/*
var migrations embed.FS

func main() {
	// Connect and migrate the DB before startup
	dbUser := env.GetString("DATABASE_USER", "")
	dbPass := env.GetString("DATABASE_PASS", "")
	dbHost := env.GetString("DATABASE_HOST", "")
	sslMode := env.GetString("SSLMODE", "require")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s?sslmode=%s", dbUser, dbPass, dbHost, sslMode)
	db, err := retryConnect(dbURL, 5)
	if err != nil {
		log.Fatalf("Failed to connect to the Database: %s", err.Error())
	}

	log.Println("Running DB migrations...")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	migrationsFS, err := iofs.New(migrations, "migrations")
	m, err := migrate.NewWithInstance("iofs", migrationsFS, "postgres", driver)
	if err != nil {
		log.Fatalf("Failed to run migrations: %s", err)
	}

	m.Up()

	s := server.New(db, Version)
	srv := &http.Server{
		ReadTimeout:  2 * time.Second,   // Time to read the request
		WriteTimeout: 10 * time.Second,  // Time to write a response
		IdleTimeout:  120 * time.Second, // Max time for keep-alive waits
		Addr:         fmt.Sprintf("%s:%s", Host, Port),
		Handler:      s,
	}

	log.Fatal(srv.ListenAndServe())
}

func retryConnect(connStr string, maxRetries int) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)
	for i := 1; i <= maxRetries; i++ {
		db, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			return db, nil
		}

		// Double sleep each time to ramp retries exponentially
		nextSleep := i * 2
		log.Printf("DB not connected: %s. Retrying in %d seconds...", err.Error(), nextSleep)
		time.Sleep(time.Duration(nextSleep) * time.Second)
	}

	return nil, fmt.Errorf("Failed to connect to the DB after %d retries", maxRetries)
}
