package networks

import "math/big"

var BSC_MAINNET Name = "bsc"

var bscMainnet = Network{
	Name:    BSC_MAINNET,
	RpcUrl:  "https://bsc-dataseed1.binance.org/",
	ChainId: big.NewInt(57),
	Testnet: false,
}

var BSC_TESTNET Name = "bsc-testnet"

var bscTestnet = Network{
	Name:    BSC_TESTNET,
	RpcUrl:  "https://data-seed-prebsc-1-s1.binance.org:8545/",
	ChainId: big.NewInt(97),
	Testnet: true,
}
