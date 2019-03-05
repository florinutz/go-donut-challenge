VERSION ?= $(shell git describe --tags 2> /dev/null || echo v0)

.PHONY: binary run lint test help fmt

.DEFAULT_GOAL := help

all: test binary ## run tests and builds the binary

binary: ## build binary for Linux
	./scripts/build/binary.sh

run: binary ## build and run the app.
	./bin $(DONUT_ARGS)

# todo replace with golangci-lint
lint:
	gometalinter ./...

test: ## run all tests
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

test_integration: ## run integration tests
	go test -v -run=\/integration -race -coverprofile=coverage.txt -covermode=atomic ./...

test_unit: ## run the unit tests
	go test -v -short -race -coverprofile=coverage.txt -covermode=atomic ./...

example_order: binary ## example for placing an order.
	./bin order --size 0.01 --price 0.100 --side buy --product-id ETH-BTC

example_ticker: binary ## example for the ticker command
	./bin ticker ETH-EUR

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
