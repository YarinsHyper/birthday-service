# Basic go commands
PROTOC=protoc

# Binary names
BINARY_NAME=birthday-service

build: build-app 
run: build
		./$(BINARY_NAME)
		MONGO_URL="mongodb://root:example@0.0.0.0:27017" ./$(BINARY_NAME)
build-app:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o $(BINARY_NAME) -v
