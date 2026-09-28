.PHONY: run build test test-race vet web-install web-dev web-test web-build

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

web-install:
	cd web && npm ci

web-dev:
	cd web && npm run dev

web-test:
	cd web && npm test

web-build:
	cd web && npm run build
