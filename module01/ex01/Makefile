# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=convert
BINARY_UNIX=$(BINARY_NAME)_unix

# file name
COVERAGE=coverage.out

# Make rule
.PHONY: all
all: build

.PHONY: build
build:
	$(GOBUILD)

.PHONY: test
test: build
	go test -run "" -race ./...

.PHONY: tests
tests:
	gotests -w -parallel -all ./...

.PHONY: cover
cover:
	$(GOTEST) ./... -cover -coverprofile=$(COVERAGE)
	go tool cover -func=$(COVERAGE)
	go tool cover -html=$(COVERAGE)

.PHONY: fuzz
fuzz:
	go test -fuzz=.

.PHONY: clean
clean:
		$(GOCLEAN)
		rm -f $(COVERAGE)

.PHONY: re
re: clean all

.PHONY: run
run: build
	./$(BINARY_NAME) $(DIR)

.PHONY: doc
doc:
	godoc -http=:6060

.PHONY: imports
imports:
	goimports -w main.go conversion/*


#.PHONY: deps
#deps:
#		$(GOGET) github.com/markbates/goth
#		$(GOGET) github.com/markbates/pop

# Cross compilation
.PHONY: build-linux
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
#.PHONY: docker-build
#docker-build:
#		docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v
