
.PHONY: build
build:
	mkdir -p grpc/
	protoc --proto_path=proto --go_out=plugins=grpc:grpc --go_opt=paths=source_relative proto/northpole.proto

.PHONY: test
test:
	 go test -v ./server
	 go test -v ./match
	 go test -v ./storage
	 go test -v ./usecase

.PHONY: testMatch
testMatch:
	 go test -v ./match
