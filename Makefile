PROJECTDIR := $(shell pwd)
GOBASE := $(HOME)/go
GOBIN=$(GOBASE)/bin
GOPATH := $(PROJECTDIR)/vendor:$(GOBASE)
TESTEXAMPLE7 := $(PROJECTDIR)/courses/task_seven

start-lint:
	@echo "  >  Start work golangci-lint"
	$(GOBIN)/golangci-lint run

# example only for lesson 7 tests, because only each subdirictory have main file
start-test:
	@echo "  >  Start work test for task_seven"
	go test -v -timeout 30s $(TESTEXAMPLE7)

start-test-coverage:
	@echo "  > Calculation code coverage for task_seven"
	# @echo $(1)
	go test -cover -v $(TESTEXAMPLE7)

# TODO take parameters from the command line
# start-dlv:
# 	@echo "  >  Start work dlv"
# 	$(GOBIN)/dlv debug $(TESTEXAMPLE7)