# App Base

A starting point for Go APIs with a Postgres DB, Redis cache, and typical authentication bits.

## Test

`make test` to run all tests and generate a coverage report as a HTML report

## Containers

A Dockerfile is provided to produce a size-optimized alpine image for deployment. Migrations are embedded
in the binary as an io.FS.
