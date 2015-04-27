# TEST_MATCHING_PATTERN="*/__test__/*test.go"

tests:
	# @find . -type f -wholename $(TEST_MATCHING_PATTERN) -print0 |xargs -0 go test -v -x
	@go test -v ./...

.PHONY: tests
