BUILD_DIR=bin
APP=app

.PHONY: build
build:
	mkdir -p ${BUILD_DIR}
	go build -o ${BUILD_DIR}/${APP} main.go

.PHONY: test
test:
	go test ./... -coverprofile out.prof
	go tool cover -html=out.prof
	rm ./out.prof

.PHONY: docs
docs:
	go get golang.org/x/tools/cmd/godoc
	@echo "Docs opening at http://localhost:6060"
	godoc -http=:6060
