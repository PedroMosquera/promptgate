.PHONY: go-run go-build go-test go-test-race go-vet web-install web-dev web-test web-build py-install py-run py-test

PY := python/.venv/bin/python

go-run:
	cd go && go run ./cmd/promptgate

go-build:
	cd go && go build -o bin/promptgate ./cmd/promptgate

go-test:
	cd go && go test ./...

go-test-race:
	cd go && go test -race ./...

go-vet:
	cd go && go vet ./...

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
	$(PY) -m pip install -r python/requirements.txt

py-run:
	cd python && PORT=8081 .venv/bin/python server.py

py-test:
	cd python && .venv/bin/python -m pytest
