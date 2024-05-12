# Go plugins for the protocol compiler:

Install the protocol compiler plugins for Go using the following commands:

- $ ```go install google.golang.org/protobuf/cmd/protoc-gen-go@latest```
- $ ```go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest```

Update your PATH so that the protoc compiler can find the plugins:

- $ ```export PATH="$PATH:$(go env GOPATH)/bin"```