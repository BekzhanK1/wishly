# Project vars
APP_NAME = wishly
PORT     = 8000

# Binary
build:
	go build -o $(APP_NAME) ./cmd/app/main.go

run:
	go run ./cmd/app/main.go

dev:
	go run ./cmd/app/main.go

test:
	go test ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

clean:
	rm -f $(APP_NAME)

# Docker
# docker-build:
# 	docker build -t $(APP_NAME) .

# docker-run:
# 	docker run -p $(PORT):$(PORT) $(APP_NAME)

.PHONY: build run dev test lint fmt clean docker-build docker-run
