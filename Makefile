BINARY_NAME=nebula-bridge

.PHONY: build run test tidy docker-build docker-run

build:
	go build -o ./bin/$(BINARY_NAME) ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./... -v

tidy:
	go mod tidy

docker-build:
	docker build -t nebula-bridge:local .

docker-run:
	docker run --rm -p 8080:8080 --env-file .env nebula-bridge:local
