.PHONY: run build test test-race vet

run:
	go run ./cmd/promptgate

build:
	go build -o bin/promptgate ./cmd/promptgate

test:
	go test ./...

test-race:
	go test -race ./...

vet:
	go vet ./...
