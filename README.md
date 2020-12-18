# northpole

## setup

```bash
# mac
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go

# document generate
go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
protoc --doc_out=html,index.html:./ proto/*.proto
```

## build 

```bash
make build
```

## format

```bash
brew install clang-format
clang-format -i proto/northpole.proto
```

## run


## reference

* https://developers.google.com/protocol-buffers/docs/proto3
