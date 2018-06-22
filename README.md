# go-loom [![Build Status](https://travis-ci.org/loomnetwork/go-loom.svg?branch=master)](https://travis-ci.org/loomnetwork/go-loom)

Go package for building Go Smart Contracts  for the Loom SDK

This package is also used for building Clients to DAppChains in the Loom SDK.

The code that runs the actual DAppChain(sidechain) is in a different repository.

## Requirements

- Go 1.9+
- Mac or Linux (Windows support coming in June)

## Installation

```bash
go get github.com/loomnetwork/go-loom
```

## Examples

The example smart contracts can be built with:
```shell
make deps
make
```
If you want the ethereum examples, use
```shell
make evm
```
instead of `make`. However you need the
[go-ethereum package](https://github.com/ethereum/go-ethereum).


To run the blockchain with the Samples

*Note Loom binary is only available to beta testers right now*

```shell
# init the blockchain
./loom init
# Copy over example genesis
cp genesis.example.json genesis.json
# run the node
./loom run
```

## Development

1. `go get` or clone the repo into your desired `GOPATH`.
2. Install deps
   ```shell
   make deps
   ```

### Generating protobufs
```shell
make proto
```

### running tests
```shell
make test
```
