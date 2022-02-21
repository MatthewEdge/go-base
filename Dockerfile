FROM golang:alpine as builder
WORKDIR /usr/src/app

# git for go mod. ca-certificates for SSL capability
RUN apk update && apk add --no-cache \
    git \
    ca-certificates \
    tzdata && \
  update-ca-certificates

COPY go.mod /usr/src/app/go.mod
RUN go mod download

COPY . /usr/src/app

# Build optimized for Scratch image
# ldflags omit symbol table, debug information, and DWARF table
ARG VERSION="0.0.0"
RUN GOOS=linux CGO_ENABLED=0 \
    go build -a -ldflags="-w -s -X main.Version=${VERSION} -X main.Host=0.0.0.0 -X main.Port=8080" \
    -o run \
    main.go

# Final, clean image
FROM alpine
WORKDIR /usr/src/app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /usr/src/app/run /usr/src/app/run

EXPOSE 8080
CMD /usr/src/app/run
