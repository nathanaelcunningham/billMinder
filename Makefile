# Load env from file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: build
build:
	go mod verify
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./build/api ./cmd/api

.PHONY: run
run:
	air

.PHONY: templ
templ:
	templ generate
