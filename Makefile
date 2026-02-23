#.PHONY: run run-movie run-metadata run-rating
reset:
	docker kill dev-consul && docker rm dev-consul

dev:
	docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
db:
	docker run --name movieexample_db -e POSTGRES_PASSWORD=password -e POSTGRES_DATABASE=movieexample -p 3306:3306 -d postgres:latest-d mysql:latest


