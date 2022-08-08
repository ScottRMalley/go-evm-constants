# Golang EVM Constants

This package is designed to set sensible defaults for popular EVM chains. It holds connection information and useful
contract addresses for the following chains:

* Avalanche
* Polygon
* Binance Smart Chain
* Ethereum

## Networks

Using network constants:

```go
package main

import (
	"github.com/scottrmalley/go-evm-constants/networks"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	avaxTestnet, err := networks.GetNetwork(networks.AVALANCHE_FUJI)
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
	customNetworkName := networks.Name("custom")
	customNetwork := networks.Network{
		Name:    customNetworkName,
		RpcUrl:  "https://network.custom.rpc/",
		ChainId: big.NewInt(1234),
		Testnet: false,
	}

	if err := networks.RegisterNetwork(customNetwork); err != nil {
		panic(err)
	}
	// this should now work
	customNetwork, err := networks.GetNetwork(customNetworkName)
	...
}
```

## Contracts

Currently, this repo holds mainnet and testnet addresses for the following DEX's:

* Pangolin
* Sushiswap
* TraderJoe

This will be continually expanded with more relevant or important contract addresses.

Getting contract addresses:
```go
package main

import (
	"github.com/scottrmalley/go-evm-constants/contracts/dex"
	"github.com/scottrmalley/go-evm-constants/networks"
)

func main() {
	contracts, err := dex.GetContracts(dex.SUSHISWAP, networks.AVALANCHE_MAINNET)
	if err != nil {
		panic(err)
    }
	...
}
```