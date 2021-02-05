.PHONY: build
build:
	mkdir -p grpc/
	protoc --proto_path=proto --go_out=plugins=grpc:grpc --go_opt=paths=source_relative proto/northpole.proto

.PHONY: docs
docs:
	protoc --doc_out=html,index.html:./ proto/*.proto
	open index.html

.PHONY: test
test:
	 go test -v ./...

