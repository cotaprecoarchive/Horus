build: tests fetch-dependencies
	@./build.sh

build-on-container:
	@docker run -it --rm -v `pwd`:/app golang:1.4.2 /bin/bash /app/container-build.sh

fetch-dependencies:
	@go get ./...

tests:
	@go test ./...

.PHONY: tests build
