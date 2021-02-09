export VERSION=$$(git describe --tags --abbrev=0)
GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build -ldflags "-X main.Version=$(VERSION)"
GOCLEAN=$(GOCMD) clean
GORUN =$(GOCMD) run main.go
GOTIDY=$(GOCMD) mod tidy
BINARY_NAME=nox
RELEASECMD=goreleaser release

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
next:
	$(RELEASECMD) --snapshot --skip-publish --rm-dist
release-dry-run:
	$(RELEASECMD)  --skip-publish --rm-dist
release:
	$(RELEASECMD)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
