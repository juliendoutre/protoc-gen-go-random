# protoc-gen-go-random

## Getting started

```shell
brew tap juliendoutre/protoc-gen-go-random https://github.com/juliendoutre/protoc-gen-go-random
brew install protoc-gen-go-random
```

## Usage

```shell
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc ./example/api.proto --go_out=./example/ --go-grpc_out=./example/ --go-random_out=./example/
```

## Run the example

```shell
cd example && go run .
# In another shell
grpcurl -plaintext localhost:8000 HelloWorld/Greet
```

## Development

### Lint the code

```shell
brew install golangci-lint
golangci-lint run
```

### Release a new version

```shell
git tag -a v0.1.0 -m "New release"
git push origin v0.1.0
```

### Update the example generated code

```shell
go install .
protoc ./example/api.proto --go_out=./example/ --go-grpc_out=./example/ --go-random_out=./example/
```

## References

- https://clement-jean.github.io/writing_protoc_plugins/
- https://github.com/theluckiestsoul/protoc-gen-structtag/blob/master/cmd/protoc-gen-structtag/main.go
