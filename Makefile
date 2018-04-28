PKG = github.com/loomnetwork/go-loom
PROTOC = protoc --plugin=./protoc-gen-gogo -Ivendor -I$(GOPATH)/src -I/usr/local/include

.PHONY: all clean test lint deps proto example-plugins example-plugins-external example-cmds

examples: example-plugins example-plugins-external example-cmds

example-cmds: create-tx

create-tx: examples/types/types.pb.go
	go build $(PKG)/examples/cmd-plugins/$@

example-plugins: helloworld.so.1.0.0

example-plugins-external: helloworld.1.0.0

helloworld.1.0.0: proto
	go build -o contracts/$@ $(PKG)/examples/plugins/helloworld

helloworld.so.1.0.0: proto
	go build -buildmode=plugin -o contracts/$@ $(PKG)/examples/plugins/helloworld

protoc-gen-gogo:
	go build github.com/gogo/protobuf/protoc-gen-gogo

%.pb.go: %.proto protoc-gen-gogo
	$(PROTOC) --gogo_out=\
plugins=grpc:$(GOPATH)/src \
$(PKG)/$<

proto: types/types.pb.go testdata/test.pb.go examples/types/types.pb.go

test: proto
	go test $(PKG)/...

lint:
	golint ./...

deps:
	go get \
		golang.org/x/crypto/ripemd160 \
		golang.org/x/crypto/sha3 \
		github.com/gogo/protobuf/jsonpb \
		github.com/gogo/protobuf/proto \
		google.golang.org/grpc \
		github.com/spf13/cobra

clean:
	go clean
	rm -f \
		protoc-gen-gogo \
		types/types.pb.go \
		testdata/test.pb.go \
		examples/types/types.pb.go \
		contracts/helloworld.1.0.0 \
		contracts/helloworld.so.1.0.0 \
		create-tx
