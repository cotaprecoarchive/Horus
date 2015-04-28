build: tests fetch-dependencies
	@./build.sh

fetch-dependencies:
	@go get ./...

tests:
	@go test -v ./...

.PHONY: tests build
