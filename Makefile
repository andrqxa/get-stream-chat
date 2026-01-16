.PHONY: run build clean docker-run test test-race lint

APP_NAME=server

run:
	go run ./cmd/server

build:
	go build -o $(APP_NAME) ./cmd/server

clean:
	rm -f $(APP_NAME)

docker-run:
	docker compose up --build --remove-orphans

test:
	go test ./...

test-race:
	go test -race ./...

lint:
	golangci-lint run
