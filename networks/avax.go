package networks

import "math/big"

var AVALANCHE_MAINNET Name = "avax"

var avalancheMainnet = Network{
	Name:    AVALANCHE_MAINNET,
	RpcUrl:  "https://api.avax.network/ext/bc/C/rpc",
	ChainId: big.NewInt(43114),
	Testnet: false,
}

var AVALANCHE_FUJI Name = "avax-fuji"

var avalancheFuji = Network{
	Name:    AVALANCHE_FUJI,
	RpcUrl:  "https://api.avax-test.network/ext/bc/C/rpc",
	ChainId: big.NewInt(43113),
	Testnet: true,
}
