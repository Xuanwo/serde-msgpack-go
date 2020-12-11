SHELL := /bin/bash

build: generate tidy
	go build ./...

generate:
	go generate ./...
	go fmt ./...

tidy:
	go mod tidy && go mod verify