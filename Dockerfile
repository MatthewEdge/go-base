FROM golang:alpine as builder
WORKDIR /usr/src/app

# git for go mod. ca-certificates for SSL capability
RUN apk update && apk add --no-cache \
    git \
    ca-certificates && \
  update-ca-certificates

COPY go.mod /usr/src/app/go.mod
RUN go mod download

COPY . /usr/src/app

# Build optimized for Scratch image
# ldflags omit symbol table, debug information, and DWARF table
RUN GOOS=linux CGO_ENABLED=0 \
    go build -a -ldflags="-w -s" -o run main.go

# Final, clean image
FROM scratch
WORKDIR /go/bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/src/app/run /go/bin/run

EXPOSE 8080

ENV RELEASE_VERSION
CMD /go/bin/run -host "0.0.0.0" -port "8080" -version "${RELEASE_VERSION}"
