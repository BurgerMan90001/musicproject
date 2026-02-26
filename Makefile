#.PHONY: run run-movie run-metadata run-rating

build-server:
	GOOS=linux go build -o ./cmd/main ./cmd/*.go

run-server:
	cd cmd && go run *.go

test:
	go test ./...