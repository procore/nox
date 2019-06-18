GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN =$(GOCMD) run main.go
GOTIDY=$(GOCMD) mod tidy
BINARY_NAME=nox

all: install
install:
	$(GOINSTALL)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
clean: 
	$(GOCLEAN)
	$(GOTIDY)
	rm -f $(BINARY_NAME)
	rm -rf ./dist
doc:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME) meta docs
	rm -f $(BINARY_NAME)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
