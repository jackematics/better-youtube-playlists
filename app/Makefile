BINARY_NAME := bootstrap 

build-LambdaFunction:
	GOOS=linux GOARCH=arm64 go build -o ./$(BINARY_NAME) main.go

clean:
	rm -f $(BINARY_NAME) 

run: build-LambdaFunction
	sam local start-api

.PHONY: build-LambdaFunction clean run