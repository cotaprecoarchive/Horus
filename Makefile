build: tests fetch-dependencies
	@./build.sh

fetch-dependencies:
	@go get ./...

tests:
	@go test ./...

.PHONY: tests build
