.PHONY: run lint
run:
	@cd sample-project && go run .

lint:
	golangci-lint run
