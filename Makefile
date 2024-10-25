BINARY_NAME := better-youtube-playlists

build:
	cd src && GOOS=linux GOARCH=amd64 go build -o ../$(BINARY_NAME)
	chmod +x $(BINARY_NAME)

run:
	sam local start-api


clean:
	rm -f $(BINARY_NAME) 

test:
	cd src && go test -v ./test/...

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
