# Use bash syntax
SHELL=/bin/bash
# Go parameters
GOCMD=go
ORG=snyk
PROJECT=driftctl
GOBINPATH=$(shell $(GOCMD) env GOPATH)/bin
GOMOD=$(GOCMD) mod
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOBINPATH)/gotestsum
MOCKERY=$(GOBINPATH)/mockery
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOTOOL=$(GOCMD) tool
GOFMT=$(GOCMD) fmt
GOOS ?= $(shell uname)
GOARCH ?= $(shell uname -m)
# ACC tests params
ACC_PATTERN ?= TestAcc_

# Dependency versions
GOTESTSUM_VERSION=v1.6.3
MOCKERY_VERSION=2.12.2

define HELP

$(ORG)/$(PROJECT) make targets
----------------------------------------------------
- help:                     Prints this
- all:                      Runs (fmt lint test build go.mod)
- build:                    Builds the binary
- test:                     Runs unit tests
- coverage:                 Generates unit test coverage report
- acc:                      Runs acceptance tests
- mocks:                    Runs acceptance tests

## Release targets

- release:       Releases the package to GitHub.
- snapshot:      Builds a release snapshot and saves the artifacts in the dist
                 folders.
- changelog:     Generates the CHANGELOG.md and notes/$(VERSION).md

endef
export HELP

.DEFAULT: help
.PHONY: help
help:
	@ echo "$$HELP"

.PHONY: FORCE

.PHONY: all
all: fmt lint test build go.mod

.PHONY: build
build:
	SINGLE_TARGET=true ./scripts/build.sh

.PHONY: release
release:
	ENV=release ./scripts/build.sh

bin/test-deps:
	go install gotest.tools/gotestsum@${GOTESTSUM_VERSION}

.PHONY: test
test: install-tools
	$(GOTEST) --format testname --junitfile unit-tests.xml -- -mod=readonly -coverprofile=cover.out.tmp -coverpkg=.,./pkg/... ./...
	cat cover.out.tmp | grep -v "mock_" > cover.out

.PHONY: coverage
coverage: test
	$(GOTOOL) cover -func=cover.out

.PHONY: acc
acc: install-tools
	DRIFTCTL_ACC=true $(GOTEST) --format testname --junitfile unit-tests-acc.xml -- -coverprofile=cover-acc.out -test.timeout 5h -coverpkg=./pkg/... -run=$(ACC_PATTERN) ./pkg/...

.PHONY: mocks
mocks: install-tools
	rm -rf mocks
	$(MOCKERY) --all

.PHONY: fmt
fmt:
	$(GOFMT) ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f bin/*

.PHONY: lint
lint:
	@which golangci-lint > /dev/null 2>&1 || (curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- -b $(GOBINPATH) v1.31.0)
	golangci-lint run -v --timeout=10m

.PHONY: install-tools
install-tools:
	$(GOINSTALL) gotest.tools/gotestsum@$(GOTESTSUM_VERSION)
	curl -L https://github.com/vektra/mockery/releases/download/v$(MOCKERY_VERSION)/mockery_$(MOCKERY_VERSION)_$(GOOS)_$(GOARCH).tar.gz -o mockery_$(MOCKERY_VERSION)_$(GOOS)_$(GOARCH).tar.gz
	tar -xjf mockery_$(MOCKERY_VERSION)_$(GOOS)_$(GOARCH).tar.gz -C $(GOBINPATH)



go.mod: FORCE
	$(GOMOD) tidy
	$(GOMOD) verify
go.sum: go.mod
