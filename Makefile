.PHONY: run install-linter lint
run:
	@cd sample-project && go run .

install-linter:
	@./install-linter

lint:
	golangci-lint run

test:
	go test ./...
