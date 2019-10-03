PKG      := github.com/mdevilliers/go
TOOLS_DIR ?= ./tools
SUBDIRS := $(wildcard */.)

# Linting
OS := $(shell uname)
GOLANGCI_LINT_VERSION=1.18.0
ifeq ($(OS),Darwin)
	GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-darwin-amd64.tar.gz
else
	GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-linux-amd64.tar.gz
endif

.PHONY: all $(SUBDIRS)
all: $(SUBDIRS)
$(SUBDIRS):  $(TOOLS_DIR)/golangci-lint/golangci-lint
ifeq ("$(wildcard $(shell which gocov))","")
	go get github.com/axw/gocov/gocov
endif
	cd $@ && gocov test ./... | gocov report
	cd $@ && ../$(TOOLS_DIR)/golangci-lint/golangci-lint run

$(TOOLS_DIR)/golangci-lint/golangci-lint:
	curl -OL https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/$(GOLANGCI_LINT_ARCHIVE)
	mkdir -p $(TOOLS_DIR)/golangci-lint/
	tar -xf $(GOLANGCI_LINT_ARCHIVE) --strip-components=1 -C $(TOOLS_DIR)/golangci-lint/
	chmod +x $(TOOLS_DIR)/golangci-lint
	rm -f $(GOLANGCI_LINT_ARCHIVE)
