# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=convert
BINARY_UNIX=$(BINARY_NAME)_unix

# Make rule
.PHONY: all
all: build

.PHONY: build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v

.PHONY: test
test: build
		./test.sh
#		$(GOTEST) -v ./...

.PHONY: clean
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)

.PHONY: fclean
fclean: clean
		find ./images -name "*.png" | xargs rm -v

.PHONY: run
run: build
	./$(BINARY_NAME)

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
