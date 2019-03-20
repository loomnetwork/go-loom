PKG = github.com/loomnetwork/go-loom
PROTOC = protoc --plugin=./protoc-gen-gogo -I$(GOPATH)/src -I/usr/local/include
GOGO_PROTOBUF_DIR = $(GOPATH)/src/github.com/gogo/protobuf
HASHICORP_DIR = $(GOPATH)/src/github.com/hashicorp/go-plugin
GETH_DIR = $(GOPATH)/src/github.com/ethereum/go-ethereum
SSHA3_DIR = $(GOPATH)/src/github.com/miguelmota/go-solidity-sha3
# This commit sha should match the one in loomchain repo
GETH_GIT_REV = 1fb6138d017a4309105d91f187c126cf979c93f9

.PHONY: all evm examples example-cli evmexample-cli example-plugins example-plugins-external plugins proto test lint deps clean test-evm deps-evm deps-all

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
	builtin/types/address_mapper/address_mapper.pb.go \
	builtin/types/coin/coin.pb.go \
	builtin/types/ethcoin/ethcoin.pb.go \
	builtin/types/dpos/dpos.pb.go \
	builtin/types/dposv2/dposv2.pb.go \
	builtin/types/dposv3/dposv3.pb.go \
	builtin/types/plasma_cash/plasma_cash.pb.go \
	builtin/types/karma/karma.pb.go \
	builtin/types/chainconfig/chainconfig.pb.go \
	builtin/types/transfer_gateway/transfer_gateway.pb.go \
	builtin/types/transfer_gateway/v1/transfer_gateway.pb.go \
	testdata/test.pb.go \
	examples/types/types.pb.go \
	examples/plugins/lottery/lottery.pb.go \
	examples/plugins/evmexample/types/types.pb.go \
	examples/plugins/evmproxy/types/types.pb.go

test: proto
	go test -v $(PKG)/...

test-evm: proto
	go test -tags "evm" -v $(PKG)/...

lint:
	golint ./...

$(SSHA3_DIR):
	git clone -q https://github.com/loomnetwork/go-solidity-sha3.git $@

$(GETH_DIR):
	git clone -q https://github.com/loomnetwork/go-ethereum.git $@

deps-all: deps deps-evm

deps:
	go get \
		golang.org/x/crypto/ripemd160 \
		golang.org/x/crypto/sha3 \
		github.com/gogo/protobuf/jsonpb \
		github.com/gogo/protobuf/proto \
		github.com/gorilla/websocket \
		github.com/phonkee/go-pubsub \
		google.golang.org/grpc \
		github.com/spf13/cobra \
		github.com/hashicorp/go-plugin \
		github.com/stretchr/testify/assert \
		github.com/go-kit/kit/log \
		github.com/pkg/errors
	dep ensure -vendor-only
	cd $(GOGO_PROTOBUF_DIR) && git checkout v1.1.1
	cd $(HASHICORP_DIR) && git checkout f4c3476bd38585f9ec669d10ed1686abd52b9961

deps-evm: $(SSHA3_DIR) $(GETH_DIR)
	cd $(GETH_DIR) && git checkout master && git pull && git checkout $(GETH_GIT_REV)
	go get \
		github.com/loomnetwork/yubihsm-go \
		gopkg.in/check.v1


clean:
	go clean
	rm -f \
		protoc-gen-gogo \
		types/types.pb.go \
		auth/auth.pb.go \
		vm/vm.pb.go \
		builtin/types/coin/coin.pb.go \
		builtin/types/ethcoin/ethcoin.pb.go \
		builtin/types/plasma_cash/plasma_cash.pb.go \
		builtin/types/karma/karma.pb.go \
		builtin/types/chainconfig/chainconfig.pb.go \
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
