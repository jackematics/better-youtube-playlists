BINARY_NAME := better-youtube-playlists 
ARTIFACTS_DIRECTORY := .aws-sam/build/LambdaFunction

build:
	cd app && GOOS=linux GOARCH=arm64 go build -o ../$(BINARY_NAME) main.go

build-local:
	cd app && go build -o $(BINARY_NAME) && HTTP_PORT=8000 $(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME) 

test:
	cd app && go test -v ./test/...

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

deploy-lambda:
	aws cloudformation deploy \
			--template-file infrastructure/better-youtube-playlists-template.json \
			--stack-name  BetterYoutubePlaylistsLambdaStack \
			--capabilities CAPABILITY_NAMED_IAM

delete-stacks:
	aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsLambdaStack
	aws cloudformation delete-stack --stack-name BetterYoutubePlaylistsS3Stack


.PHONY: build test clean 
