.PHONY: run lint
run:
	@cd sample-project && go run .

lint:
	golangci-lint run

test:
	go test ./...

build:
	go build -o fatcontext-linter ./cmd/fatcontext
