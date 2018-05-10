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

## Install
After building wrapstore, copy `wrapstore.1.0.0` and `SimpleStore.json` to the loom contracts directory.
After initialising the DAppChain `./loom init`, add the `genesis.json` file from here to your loomchain directory. 
Ammend the location entry in the `genesis.json` to match your system.
Now run the chain with `./loom run`

```bash
go build -o wrapstore.1.0.0  wrapstore.go
export LOOMCHAIN_DIR=\My\Loomchain\Directory
$LOOMCHAIN_DIR/loom init
cp wrapstore.1.0.0 SimpleStore.json genesis.json $LOOMCHAIN_DIR
$LOOMCHAIN_DIR/loom run
```

## Testing

You can now run the wrapstore/cli.go tool to acess the contract.
```bash
./simplestore set 3455
./simplestore get
```

