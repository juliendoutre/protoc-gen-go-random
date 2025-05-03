# protoc-gen-go-random

## Getting started

```shell
# Install dependencies
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# Install the plugin
brew tap juliendoutre/protoc-gen-go-random https://github.com/juliendoutre/protoc-gen-go-random
brew install protoc-gen-go-random
# Run the example
cd example
go run .
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
