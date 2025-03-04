BINARY_NAME := better-youtube-playlists 
OS := mac

build:
ifeq ($(OS),mac)
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME) main.go
else
	GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME) main.go
endif

clean:
	rm -f $(BINARY_NAME) 

test:
	go test -v ./test/...

run:
	$(MAKE) build
	./$(BINARY_NAME)

# database

localdb: 
	docker compose up

kill-localdb:
	docker compose kill; docker compose rm -vf

# infrastructure

deploy-s3:
	aws cloudformation deploy \
		--template-file infrastructure/s3-template.json \
		--stack-name  BetterYoutubePlaylistsS3Stack \
		--capabilities CAPABILITY_NAMED_IAM 

deploy-ec2:
	aws cloudformation deploy \
		--template-file infrastructure/ec2-template.json \
		--stack-name BetterYoutubePlaylistsEC2Stack \
		--capabilities CAPABILITY_NAMED_IAM

delete-s3-stack:
	aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsS3Stack

delete-ec2-stack:
	aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsEC2Stack


.PHONY: build test clean 
