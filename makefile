# Define variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GORUN=$(GOCMD) run
BINARY_NAME=main
SOURCE_FILE=main.go

# Default target
all: build

# Build target
build:
	$(GOBUILD) -o $(SOURCE_FILE) -v
	@echo "Build  Created succesfully" 

# Clean target
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Build and Run target
build-run:
	$(GOBUILD) -o $(SOURCE_FILE) -v
	./$(BINARY_NAME)

# Run target
run:
	$(GORUN) $(SOURCE_FILE)
	@echo "Running" 
	

# .PHONY rule to avoid conflicts with filenames
.PHONY: all build clean build-run