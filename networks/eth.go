package networks

import "math/big"

var ETHEREUM_MAINNET Name = "eth"

var ethereum = Network{
	Name:    ETHEREUM_MAINNET,
	RpcUrl:  "https://rpc.ankr.com/eth/",
	ChainId: big.NewInt(1),
	Testnet: false,
}

var ETHEREUM_GOERLI Name = "eth-goerli"

var ethereumGoerli = Network{
	Name:    ETHEREUM_GOERLI,
	RpcUrl:  "https://ethereum-goerli-rpc.allthatnode.com/",
	ChainId: big.NewInt(5),
	Testnet: true,
}
