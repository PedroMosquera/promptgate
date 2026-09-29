.PHONY: run build test test-race vet web-install web-dev web-test web-build py-install py-run py-test

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

py-install:
	python3 -m venv python/.venv
	python/.venv/bin/python3 -m pip install -r python/requirements.txt

py-run:
	cd python && PORT=8081 .venv/bin/python3 server.py

py-test:
	cd python && .venv/bin/python3 -m pytest
