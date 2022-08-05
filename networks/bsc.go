package networks

import "math/big"

var BSC Name = "bsc"

var BSC_BLOCKCHAIN = Blockchain{
	Mainnet: Network{
		Name:    BSC,
		RpcUrl:  "https://bsc-dataseed1.binance.org/",
		ChainId: big.NewInt(57),
		Testnet: false,
	},
	Testnet: Network{
		Name:    BSC,
		RpcUrl:  "https://data-seed-prebsc-1-s1.binance.org:8545/",
		ChainId: big.NewInt(97),
		Testnet: true,
	},
}
