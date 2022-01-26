GOCMD	= go
GOBUILD	= $(GOCMD) build
GOCLEAN	= $(GOCMD) clean
GOTEST	= $(GOCMD) test
GOGET	= $(GOCMD) get

.PHONY: all
all: build

.PHONY: build
build:
	$(GOBUILD)

.PHONY: clean
clean:
	$(GOCLEAN)

.PHONY: re
re: clean all
