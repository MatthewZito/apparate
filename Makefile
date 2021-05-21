GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_NAME=apparate
BIN_UNIX=$(BIN_NAME)_unix

WRAPPER=apparate.bash

all: clean test build

build: 
	$(GOBUILD) -o $(BIN_NAME) -v && chmod u+x $(WRAPPER)

.PHONY: test
test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BIN_NAME)
	rm -f $(BIN_UNIX)