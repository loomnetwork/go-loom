## Wrapstore

This example plugin shows how to call a solidity contract running on the DAppChain's virtual machine from a loom plugin.

It wraps the SimpleStore contract.
```solidity
pragma solidity ^0.4.18;
contract SimpleStore {
  function set(uint _value) public {
    value = _value;
  }

  function get() public constant returns (uint) {
    return value;
  }

  uint value;
}
```
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

Build the wrapstore library.
Assuming loom is in your path initialise the DAppChain `./loom init`. 
copy the `example.genesis.json` to `genesis.json` and then 
run the chain with `./loom run`

```bash
go build -tags "evm" -o contracts/wrapstore.1.0.0  wrapstore.go
loom init
cp example.genesis.json genesis.json
loom run
```

## Testing
The cli can be bult with
```bash
cd ../cli
go build simplestore.go
```

You can now run the wrapstore/cli.go tool to access the solidty contract. 
You might need to use -r and -w to the DAppChain's URL.
```bash
./simplestore set -v 3455
./simplestore get
```

