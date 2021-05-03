BIN := sbm
VERSION := $$(make -s show-version)
VERSION_PATH := .
GOBIN ?= $(shell go env GOPATH)/bin
BUILD_LDFLAGS := "-s -w -X main.revision=$(CURRENT_REVISION)"
.DEFAULT_GOAL := help

.PHONY: help
help: ## この文章を表示します。
	# http://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build
	@go build -ldflags=$(BUILD_LDFLAGS) .

.PHONY: clean
clean: ## clean up bin
	@rm -f $(BIN)

.PHONY: show-version
show-version: $(GOBIN)/gobump ## show version
	@gobump show -r $(VERSION_PATH)

.PHONY: bump
bump: ## bump up version interactively
ifneq ($(shell git status --porcelain),)
	$(error git workspace is dirty)
endif
ifneq ($(shell git rev-parse --abbrev-ref HEAD),master)
	$(error current branch is not master)
endif
	@gobump up -w "$(VERSION_PATH)"
	git commit -am "bump up version to $(VERSION)"
	git tag "v$(VERSION)"
	git push origin master
	git push origin "refs/tags/v$(VERSION)"
