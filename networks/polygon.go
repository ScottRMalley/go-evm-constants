package networks

import "math/big"

var POLYGON Name = "matic"

var POLYGON_BLOCKCHAIN = Blockchain{
	Mainnet: Network{
		Name:    POLYGON,
		RpcUrl:  "https://polygon-rpc.com/",
		ChainId: big.NewInt(147),
		Testnet: false,
	},
	Testnet: Network{
		Name:    POLYGON,
		RpcUrl:  "https://rpc-mumbai.matic.today/",
		ChainId: big.NewInt(80001),
		Testnet: true,
	},
}
