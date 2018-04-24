# loom-plugin
Go package for building plugins for the Loom SDK

## Requirements

- Go 1.9+
- Mac or Linux (Windows is not supported)

## Installation

```bash
go get github.com/loomnetwork/loom-plugin
# dependencies
go get github.com/spf13/cobra golang.org/x/crypto github.com/gogo/protobuf
```

## Examples

The example plugins can be built with:

```shell
go build -buildmode=plugin -o out/cmds/create-tx.so examples/cmd-plugins/create-tx/main.go
```

## Development

1. `go get` or clone the repo into your desired `GOPATH`.
2. Ensure you have `https://github.com/golang/dep` installed, and then run the
   following command to fetch the dependencies:

   ```shell
   dep ensure
   ```

### Generating protobufs
```shell
go build github.com/gogo/protobuf/protoc-gen-gogo
protoc --plugin=./protoc-gen-gogo -Ivendor -I$GOPATH/src --gogo_out=\
plugins=grpc,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:$GOPATH/src github.com/loomnetwork/loom-plugin/types/types.proto
```
