.PHONY: example-simple
example-simple:
	@go run ./examples/simple/*.go

.PHONY: example-viewport
example-viewport:
	@go run ./examples/viewport/*.go

.PHONY: test
test:
	@go test -race -cover ./datatree

.PHONY: test-coverage
test-coverage: coverage.out
	@go tool cover -html=coverage.out

.PHONY: benchmark
benchmark:
	@go test -run=XXX -bench=. -benchmem ./datatree

.PHONY: lint
lint: ./bin/golangci-lint
	@./bin/golangci-lint run ./datatree

coverage.out: datatree/*.go go.*
	@go test -coverprofile=coverage.out ./datatree

.PHONY: fmt
fmt:
	@go fmt ./...

./bin/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.44.2

