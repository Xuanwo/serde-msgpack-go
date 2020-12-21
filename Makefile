SHELL := /bin/bash

build: generate tidy
	go build ./...

generate:
	go generate ./...
	go fmt ./...

test:
	go test -v ./...

tidy:
	go mod tidy && go mod verify