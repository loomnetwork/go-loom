module github.com/loomnetwork/go-loom

go 1.13

require (
	github.com/aristanetworks/goarista v0.0.0-20200129185426-4481d91782e5 // indirect
	github.com/btcsuite/btcd v0.0.0-20190109040709-5bda5314ca95
	github.com/certusone/yubihsm-go v0.1.1-0.20190814054144-892fb9b370f3
	github.com/cespare/cp v1.1.1 // indirect
	github.com/cevaris/ordered_map v0.0.0-20190319150403-3adeae072e73 // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/ethereum/go-ethereum v1.9.10
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5 // indirect
	github.com/go-kit/kit v0.9.0
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/gogo/protobuf v1.1.1
	github.com/gorilla/websocket v1.4.1
	github.com/grpc-ecosystem/go-grpc-prometheus v0.0.0-20180418170936-39de4380c2e0
	github.com/hashicorp/go-plugin v0.0.0-20181211201406-f4c3476bd385
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/huin/goupnp v1.0.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/karalabe/hid v1.0.0 // indirect
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/kisielk/gotool v1.0.0 // indirect
	github.com/loomnetwork/mamamerkle v0.0.0-20200206113614-cc12f6675a88 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/miguelmota/go-solidity-sha3 v0.1.0
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/phonkee/go-pubsub v0.0.0-20181130135233-5425e5981d13
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20190404164418-38d8ce5564a5
	golang.org/x/net v0.0.0-20190912160710-24e19bdeb0f2
	golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9 // indirect
	google.golang.org/grpc v1.23.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
)

replace github.com/miguelmota/go-solidity-sha3 => github.com/loomnetwork/go-solidity-sha3 v0.0.2-0.20190227083338-45494d847b31

replace github.com/ethereum/go-ethereum => github.com/loomnetwork/go-ethereum 8e02782666c8131d6327dc01522efcddbfff9f01
