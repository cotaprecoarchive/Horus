fetch-dependencies:
	@go get ./...

tests:
	@go test -v ./...

.PHONY: tests
