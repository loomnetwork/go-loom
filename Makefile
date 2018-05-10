PKG = github.com/loomnetwork/go-loom
PROTOC = protoc --plugin=./protoc-gen-gogo -Ivendor -I$(GOPATH)/src -I/usr/local/include

.PHONY: all clean test lint deps proto examples example-plugins example-plugins-external example-cmds

all: examples

examples: example-plugins example-plugins-external example-cli

example-cli: examples/types/types.pb.go
	go build -o $@ $(PKG)/examples/cli

example-plugins: contracts/helloworld.so.1.0.0 contracts/lottery.so.1.0.0

example-plugins-external: contracts/helloworld.1.0.0

contracts/helloworld.1.0.0: proto
	go build -o $@ $(PKG)/examples/plugins/helloworld

contracts/helloworld.so.1.0.0: proto
	go build -buildmode=plugin -o $@ $(PKG)/examples/plugins/helloworld

contracts/lottery.so.1.0.0: examples/plugins/lottery/lottery.pb.go
	go build -o $@ $(PKG)/examples/plugins/lottery

protoc-gen-gogo:
	go build github.com/gogo/protobuf/protoc-gen-gogo

%.pb.go: %.proto protoc-gen-gogo
	$(PROTOC) --gogo_out=plugins=grpc:$(GOPATH)/src $(PKG)/$<

proto: \
	types/types.pb.go \
	auth/auth.pb.go \
	vm/vm.pb.go \
	plugin/types/types.pb.go \
	builtin/types/coin/coin.pb.go \
	builtin/types/dpos/dpos.pb.go \
	testdata/test.pb.go \
	examples/types/types.pb.go \
	examples/plugins/lottery/lottery.pb.go

test: proto
	go test -v $(PKG)/...

lint:
	golint ./...

deps:
	go get \
		golang.org/x/crypto/ripemd160 \
		golang.org/x/crypto/sha3 \
		github.com/gogo/protobuf/jsonpb \
		github.com/gogo/protobuf/proto \
		google.golang.org/grpc \
		github.com/spf13/cobra \
		github.com/hashicorp/go-plugin \
		github.com/stretchr/testify/assert

clean:
	go clean
	rm -f \
		protoc-gen-gogo \
		types/types.pb.go \
		auth/auth.pb.go \
		vm/vm.pb.go \
		builtin/types/coin/coin.pb.go \
		testdata/test.pb.go \
		examples/types/types.pb.go \
		builtin/plugins/lottery/lottery.pb.go \
		contracts/helloworld.1.0.0 \
		contracts/helloworld.so.1.0.0 \
		out/cmds/cli \
