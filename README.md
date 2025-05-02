# protoc-gen-go-random

## Prerequisites

```shell
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Run the example

```shell
go install .
protoc ./example/api.proto --go_out=./example/ --go-grpc_out=./example/ --go-random_out=./example/
```

## References

- https://clement-jean.github.io/writing_protoc_plugins/
- https://github.com/theluckiestsoul/protoc-gen-structtag/blob/master/cmd/protoc-gen-structtag/main.go
