.PHONY: build

build:
	go get -u ./...
	go build -o ./bin/client
	docker build --tag client .