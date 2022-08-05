package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/scottrmalley/go-evm-constants/networks"
)

var TRADER_JOE Name = "trader-joe"

// TRADER_JOE_DEX values pulled from the Trader Joe SDK:
// https://github.com/traderjoe-xyz/joe-sdk/blob/main/src/constants.ts
var TRADER_JOE_DEX = Dex{
	Name: TRADER_JOE,
	Networks: map[networks.Name]NetworkDex{
		networks.AVALANCHE: {
			Mainnet: Contracts{
				Router:  common.HexToAddress("0x60aE616a2155Ee3d9A68541Ba4544862310933d4"),
				Factory: common.HexToAddress("0x9Ad6C38BE94206cA50bb0d90783181662f0Cfa10"),
			},
			Testnet: Contracts{
				Router:  common.HexToAddress("0xd7f655E3376cE2D7A2b08fF01Eb3B1023191A901"),
				Factory: common.HexToAddress("0xF5c7d9733e5f53abCC1695820c4818C59B457C2C"),
			},
		},
	},
}
