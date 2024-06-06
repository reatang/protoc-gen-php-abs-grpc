
build:
	go build -o $(GOPATH)/bin/protoc-gen-php-abs-grpc cmd/protoc-gen-php-abs-grpc/main.go

build_test:
	go build -o protoc-gen-php-abs-grpc cmd/protoc-gen-php-abs-grpc/main.go