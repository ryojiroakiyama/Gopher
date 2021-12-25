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
	go test -v ./...

.PHONY: cover
cover:
	$(GOTEST) -v ./... -cover -coverprofile $(COVERAGE)
	go tool cover -func=$(COVERAGE)
	go tool cover -html=$(COVERAGE)

.PHONY: clean
clean:
		$(GOCLEAN)
		rm $(COVERAGE)

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
	goimports -w main.go converter/*

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

#goimports #code clean # -w -> overwrite
#golint #detect unlike gostyle #not recommended
#go vet #detect error like bug
#errcheck #detect miss of error checking
#statickcheck #thirdparty
#GolangCI-Lint #thrdparty Linter Runner
