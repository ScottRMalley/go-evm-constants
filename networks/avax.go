package networks

import "math/big"

var AVALANCHE Name = "avax"

var AVALANCHE_BLOCKCHAIN = Blockchain{
	Mainnet: Network{
		Name:    AVALANCHE,
		RpcUrl:  "https://api.avax.network/ext/bc/C/rpc",
		ChainId: big.NewInt(43114),
		Testnet: false,
	},
	Testnet: Network{
		Name:    AVALANCHE,
		RpcUrl:  "https://api.avax-test.network/ext/bc/C/rpc",
		ChainId: big.NewInt(43113),
		Testnet: true,
	},
}
