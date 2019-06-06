GOCMD=go
RELEASECMD=goreleaser
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
next:
	$(RELEASECMD) --snapshot --skip-publish --rm-dist
release-dry-run:
	$(RELEASECMD) release --skip-publish --rm-dist
release:
	$(RELEASECMD)
doc:
	$(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME) run meta docs
	rm -f $(BINARY_NAME)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
