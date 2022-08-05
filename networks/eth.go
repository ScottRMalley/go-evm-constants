package networks

import "math/big"

var ETHEREUM Name = "eth"

var ETHEREUM_BLOCKCHAIN = Blockchain{
	Mainnet: Network{
		Name:    ETHEREUM,
		RpcUrl:  "https://mainnet.infura.io/v3/",
		ChainId: big.NewInt(1),
		Testnet: false,
	},
	Testnet: Network{
		Name:    ETHEREUM,
		RpcUrl:  "https://goerli.infura.io/v3/",
		ChainId: big.NewInt(5),
		Testnet: true,
	},
}
