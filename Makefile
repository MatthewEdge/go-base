BUILD_DIR=bin
APP=app

.DEFAULT_GOAL: help

COMPOSE=docker-compose

.PHONY: help
help:
	@echo "Options:\n"
	@sed -n 's|^##||p' ${PWD}/Makefile

## build: Build the app's container environment. Optional: app={SERVICE_NAME}
.PHONY: build
build:
	${COMPOSE} build ${app}

## start: Start the container environment. Optional: app={SERVICE_NAME}
.PHONY: start
start:
	${COMPOSE} up -d ${app}

## logs: Tail logs. Optional: app={SERVICE_NAME}
.PHONY: logs
logs:
	${COMPOSE} logs -f ${app}

## stop: Stop the container environment. Optional: app={SERVICE_NAME}
.PHONY: stop
stop:
	${COMPOSE} stop  ${app}

## dbcli: Connect to the DB container using pgsql
.PHONY: dbcli
dbcli:
	docker run --rm -it --network go-base_default -e PGPASSWORD=dontUseThisInProduction postgres:14-alpine psql -h db -U postgres -d appdb -x

## run: Run the app directly
.PHONY: run
run:
	go run main.go

## test: Run unit tests and open the Converage HTML report
.PHONY: test
test:
	go test ./... -coverprofile out.prof
	go tool cover -html=out.prof
	rm ./out.prof

## docs: Run the Go Docs server
.PHONY: docs
docs:
	go get golang.org/x/tools/cmd/godoc
	@echo "Docs opening at http://localhost:6060"
	godoc -http=:6060
