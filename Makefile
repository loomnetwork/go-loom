PKG = github.com/loomnetwork/loom-plugin
PROTOC = protoc --plugin=./protoc-gen-gogo -Ivendor -I$(GOPATH)/src -I/usr/local/include

.PHONY: all clean test deps proto

all: proto

protoc-gen-gogo:
	go build github.com/gogo/protobuf/protoc-gen-gogo

%.pb.go: %.proto protoc-gen-gogo
	$(PROTOC) --gogo_out=\
plugins=grpc,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:$(GOPATH)/src \
$(PKG)/$<

proto: types/types.pb.go

test: proto
	go test $(PKG)/...

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
	rm -f ./protoc-gen-gogo
	rm -f types/types.pb.go
