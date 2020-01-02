GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

REPO_NAME=repo

GITHUB_USER=devlights
PKG_NAME=mkghurl
BIN_NAME=mkghurl
CMD_PKG=github.com/$(GITHUB_USER)/$(PKG_NAME)

.PHONY: all
all: clean build test

.PHONY: build
build:
	$(GOBUILD) -o $(BIN_NAME) -v $(CMD_PKG)

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -f ./$(BIN_NAME)

.PHONY: run
run: clean build
	./$(BIN_NAME) $(REPO_NAME)
	rm -f ./$(BIN_NAME)

.PHONY: version
version: clean build
	./$(BIN_NAME) -v
	rm -f ./$(BIN_NAME)

