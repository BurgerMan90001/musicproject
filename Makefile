#.PHONY: run run-movie run-metadata run-rating

build-server:
	GOOS=linux go build -o ./config/main ./cmd/*.go

run-server:
	cd config && sh ./main

test:
	go test ./...