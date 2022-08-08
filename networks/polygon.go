package networks

import "math/big"

var POLYGON_MAINNET Name = "matic"

var polygonMainnet = Network{
	Name:    POLYGON_MAINNET,
	RpcUrl:  "https://polygon-rpc.com/",
	ChainId: big.NewInt(147),
	Testnet: false,
}

var POLYGON_TESTNET Name = "matic-testnet"

var polygonTestnet = Network{
	Name:    POLYGON_TESTNET,
	RpcUrl:  "https://rpc-mumbai.matic.today/",
	ChainId: big.NewInt(80001),
	Testnet: true,
}
