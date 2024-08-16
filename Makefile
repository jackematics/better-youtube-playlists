BINARY_NAME := better-youtube-playlists

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v ./test/...

run: build
	./$(BINARY_NAME)

all: build

localdb: 
	docker compose up -d

kill-localdb:
	docker compose kill; docker compose rm -vf

.PHONY: build test clean deps run all
