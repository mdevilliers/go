PKG      := github.com/mdevilliers/go
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
TOOLS_DIR ?= ./tools

# Linting
OS := $(shell uname)
GOLANGCI_LINT_VERSION=1.18.0
ifeq ($(OS),Darwin)
	GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-darwin-amd64.tar.gz
else
	GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64.tar.gz
endif

# the linting gods must be obeyed
.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint/golangci-lint
	$(TOOLS_DIR)/golangci-lint/golangci-lint run

$(TOOLS_DIR)/golangci-lint/golangci-lint:
	curl -OL https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/$(GOLANGCI_LINT_ARCHIVE)
	mkdir -p $(TOOLS_DIR)/golangci-lint/
	tar -xf $(GOLANGCI_LINT_ARCHIVE) --strip-components=1 -C $(TOOLS_DIR)/golangci-lint/
	chmod +x $(TOOLS_DIR)/golangci-lint
	rm -f $(GOLANGCI_LINT_ARCHIVE)

.PHONY: test
# Run test suite
test:
ifeq ("$(wildcard $(shell which gocov))","")
	go get github.com/axw/gocov/gocov
endif
	gocov test ${PKG_LIST} | gocov report

