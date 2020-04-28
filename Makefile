PROJECTDIR := $(shell pwd)
GOBASE := $(HOME)/go
GOBIN=$(GOBASE)/bin
GOPATH := $(PROJECTDIR)/vendor:$(GOBASE)

start-lint:
	@echo "  >  Start work golangci-lint"
	$(GOBIN)/golangci-lint run
