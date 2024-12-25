BINARY_NAME := better-youtube-playlists 
ARTIFACTS_DIRECTORY := .aws-sam/build/LambdaFunction

build:
	GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME) 

test:
	go test -v ./test/...

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

delete-stacks:
	aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsEC2Stack
	aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsS3Stack


.PHONY: build test clean 
