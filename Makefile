VERSION ?= $(shell git describe --tags 2> /dev/null || echo v0)

.PHONY: binary run lint test help fmt require_coinbase_vars

.DEFAULT_GOAL := help

all: test binary ## run tests and builds the binary

binary: ## build binary for Linux
	./scripts/build/binary.sh

run: binary ## build and run the app. example: COINBASE_PRO_PASSPHRASE="..." COINBASE_PRO_KEY="..." COINBASE_PRO_SECRET="..." COINBASE_PRO_SANDBOX=1 DONUT_ARGS="order --size 0.01 --price 0.100 --side buy --product-id BTC-USD" make run
	./bin $(DONUT_ARGS)

lint:
	gometalinter ./...

# same as below (require_coinbase_vars)
test: ## run all tests
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# add a dependency for require_coinbase_vars when adding order tests
test_integration: ## run integration tests (order tests require_coinbase_vars)
	go test -v -run=\/integration -race -coverprofile=coverage.txt -covermode=atomic ./...

test_unit: ## run the unit tests
	go test -v -short -race -coverprofile=coverage.txt -covermode=atomic ./...

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

require_coinbase_vars:
	help := is not set. See https://github.com/preichenberger/go-coinbasepro#setup
	ifndef COINBASE_PRO_PASSPHRASE
		$(error COINBASE_PRO_PASSPHRASE $(help))
	endif
	ifndef COINBASE_PRO_KEY
		$(error COINBASE_PRO_KEY $(help))
	endif
	ifndef COINBASE_PRO_SECRET
		$(error COINBASE_PRO_SECRET $(help))
	endif
