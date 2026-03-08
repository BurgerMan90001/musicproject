#.PHONY: run run-movie run-metadata run-rating

build:
	docker build . -t musicproject/server

#GOOS=linux go build -o ./config/main ./cmd/*.go

run:
	docker run -p 8081:8081 musicproject/server
#cd config && sh ./main

test:
	go test ./...