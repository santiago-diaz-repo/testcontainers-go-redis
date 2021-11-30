PROJECT_NAME := "testcontainers-go-redis"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

all: clean fmt lint security test build ## Execute clean, gen-mock, fmt, lint, security, test and build

build: ## Build the binary file
	go build .

clean: ## Clean dev files
	go clean -i ./...

coverage: ## Generate global code coverage report in a file called cover.html
	@go test -coverprofile cover.out.tmp ./... &&\
	cat cover.out.tmp | grep -v "_mock.go" > cover.out &&\
	go tool cover -html=cover.out -o cover.html &&\
	rm cover.out.tmp cover.out

dependencies: ## Download dependencies
	@go mod download

fmt: ## Format src code files
	@go fmt ./...

install: build ## Install the binary file
	@go install

lint: ## Execute lint
	@golangci-lint run ./...

security: ## Execute go sec to check secure code
	gosec -tests ./...

test: ## Execute tests
	@go test ${PKG_LIST}

tidy:  ## Execute tidy command
	@go mod tidy

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: \
	all \
	build \
	clean \
	coverage \
	dependencies \
	fmt \
	install \
	lint \
	security \
	test \
	tidy \