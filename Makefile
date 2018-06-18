PKG = github.com/loomnetwork/go-loom
PROTOC = protoc --plugin=./protoc-gen-gogo -I$(GOPATH)/src -I/usr/local/include

.PHONY: all evm examples example-cli evmexample-cli example-plugins example-plugins-external plugins proto test lint deps clean

all: examples

evm: all example-evm-plugins evmexample-cli

examples: example-plugins example-plugins-external example-cli

example-cli: proto
	go build -o $@ $(PKG)/examples/cli

evmexample-cli: proto
	go build -tags "evm" -o $@ $(PKG)/examples/plugins/evmexample/cli

example-plugins: contracts/helloworld.so.1.0.0 contracts/lottery.so.1.0.0

example-plugins-external: contracts/helloworld.1.0.0

example-evm-plugins: contracts/evmexample.1.0.0 contracts/evmproxy.1.0.0

contracts/helloworld.1.0.0: proto
	go build -o $@ $(PKG)/examples/plugins/helloworld

contracts/helloworld.so.1.0.0: proto
	go build -buildmode=plugin -o $@ $(PKG)/examples/plugins/helloworld

contracts/lottery.so.1.0.0: examples/plugins/lottery/lottery.pb.go
	go build -o $@ $(PKG)/examples/plugins/lottery

contracts/evmexample.1.0.0: proto
	go build -tags "evm" -o $@ $(PKG)/examples/plugins/evmexample/contract

contracts/evmproxy.1.0.0: proto
	go build -tags "evm" -o $@ $(PKG)/examples/plugins/evmproxy/contract

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
	builtin/types/plasma/plasma_cash.pb.go \
	testdata/test.pb.go \
	examples/types/types.pb.go \
	examples/plugins/lottery/lottery.pb.go \
	examples/plugins/evmexample/types/types.pb.go \
	examples/plugins/evmproxy/types/types.pb.go

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
		github.com/stretchr/testify/assert \
		github.com/go-kit/kit/log
	dep ensure -vendor-only

clean:
	go clean
	rm -f \
		protoc-gen-gogo \
		types/types.pb.go \
		auth/auth.pb.go \
		vm/vm.pb.go \
		builtin/types/coin/coin.pb.go \
		builtin/types/plasma/plasma_cash.pb.go \
		testdata/test.pb.go \
		examples/types/types.pb.go \
		examples/plugins/evmexample/types/types.pb.go \
		example-cli \
		evmexample-cli \
		builtin/plugins/lottery/lottery.pb.go \
		contracts/helloworld.1.0.0 \
		contracts/helloworld.so.1.0.0 \
		out/cmds/cli \
		contracts/evmexample.1.0.0 \
		contracts/lottery.so.1.0.0 \
		contracts/evmproxy.so.1.0.0 \
		out/cmds/cli
