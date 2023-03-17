# 'make' command will trigger the help target
.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help 
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {gsub("\\\\n",sprintf("\n%22c",""), $$2);printf "\033[36m%-20s\033[0m \t\t%s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Lint the code 
	@if [ -x "$$(command -v golangci-lint)" ]; then \
		golangci-lint run; \
	else \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.27.0; \
		golangci-lint run; \
		rm $(go env GOPATH)/bin/golangci-lint; \
	fi

