# Golang EVM Constants

This package is designed to set sensible defaults for popular EVM chains. It holds connection information and useful
contract addresses for the following chains:

* Avalanche
* Polygon
* Binance Smart Chain
* Ethereum

## Usage

Using network constants:

```go
package main

import (
	"github.com/scottrmalley/go-evm-constants/networks"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	avaxTestnet, err := networks.GetNetwork(networks.AVALANCHE, true)
	if err != nil {
		panic(err)
	}

	client, err := ethclient.Dial(avaxTestnet.RpcUrl)
	...
}
```

Adding a new network:

```golang
package main

import (
	"github.com/scottrmalley/go-evm-constants/networks"
	"math/big"
)

func main() {
	customNetworkName := NetworkName("custom")
	customBc := networks.Blockchain{
		Mainnet: Network{
			Name:    customNetworkName,
			RpcUrl:  "https://network.custom.rpc/",
			ChainId: big.NewInt(1234),
			Testnet: false,
		},
	}
	
	if err := networks.RegisterBlockchain(customNetworkName, customBc); err != nil {
		panic(err)
    }
	// this should now work
	customNetwork, err := networks.GetNetwork(customNetworkName, false)
	...
}
```

## Known Missing Elements

Currently, there is no support for multiple testnet networks on one blockchain. This is primarily relevant to 
Ethereum, as it has Goerli, Ropsten etc. For simplicity at this time, we default Eth testnet to Goerli.