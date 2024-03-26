PROJECTNAME=$(shell basename "$(PWD)")

# Go переменные.
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)

## start: Compiles, starts nginx and executes.
start:
	@echo
	@bash -c "trap 'make -s stop' EXIT; \
	$(MAKE) -s go-compile start-server exec"
	@echo

## stop: Stops nginx.
stop: stop-server

## clean: Cleans build cache and removes bin
clean: go-clean

go-compile: go-clean go-build

exec:
	@echo "  >  Executing $(PROJECTNAME)"
	@$(GOBIN)/$(PROJECTNAME)

start-server:
	@echo "  >  Starting web-server"
	@bash -c "sudo systemctl start nginx"

stop-server:
	@echo "  >  Stopping web-server"
	@bash -c "sudo nginx -s stop"

go-build:
	@echo "  >  Building binary"
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache and removing bin"
	@GOBIN=$(GOBIN) go clean
	@rm -rf bin

.PHONY: help
help: Makefile
	@echo "\n Choose a command run in "$(PROJECTNAME)":\n"
	@sed -n 's/^##/ /p' $<
	@echo 

## fmt: Formats Go programs
.PHONY: fmt
fmt:
	@echo "  >  Formatting Go programs..."
	@gofmt -w $(GOFILES)
