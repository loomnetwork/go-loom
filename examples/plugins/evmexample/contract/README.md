## evmexample

This example plugin shows how to call a solidity contract running on the DAppChain's virtual machine from a loom plugin.

It wraps the smart contract in SimpleStore.sol. If you have the solidity compiler solc installed
you can generate the abi and bin files with.
```bash
solc  --bin -o ./contracts SimpleStore.sol
solc  --abi -o. SimpleStore.sol
```
If you run the DAppChain in another directory, both `solc` output files need to be copied across.

## Prerequisite
This example requires the 
[go-ethereum package](https://github.com/ethereum/go-ethereum),
if you do not already have it installed you can use.
```bash
go get -d github.com/ethereum/go-ethereum
```

## Install
Amend the location entry in the `genesis.json` for SimpleStore to match the 
path of your go-loom directory.

`ed genesis.json`

Build the evmexample library.
Assuming loom is in your path initialise the DAppChain `./loom init`. 
copy the `example.genesis.json` to `genesis.json` and then 
run the chain with `./loom run`

```bash
go build -tags "evm" -o contracts/evmexample.1.0.0  evmexample.go
loom init
cp example.genesis.json genesis.json
loom run
```

## Testing
The cli can be bult with
```bash
cd ../cli
go build evmexample.go
```

You can now run the evmexample/cli.go tool to access the solidty contract. 
You might need to use -r and -w to set the DAppChain's URL.
```bash
./evmexample set -v 3455
./evmexample get
```

