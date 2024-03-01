BINARY_NAME := better-youtube-playlists

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v ./test/...

run: build
	./$(BINARY_NAME)

all: build

.PHONY: build test clean deps run all
