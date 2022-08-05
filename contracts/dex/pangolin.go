package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/scottrmalley/go-evm-constants/networks"
)

var PANGOLIN Name = "pangolin"

// PANGOLIN_DEX values pulled from the pangolin TS SDK:
// https://github.com/pangolindex/sdk/blob/master/src/chains.ts
var PANGOLIN_DEX = Dex{
	Name: PANGOLIN,
	Networks: map[networks.Name]NetworkDex{
		networks.AVALANCHE: {
			Mainnet: Contracts{
				Router:  common.HexToAddress("0xE54Ca86531e17Ef3616d22Ca28b0D458b6C89106"),
				Factory: common.HexToAddress("0xefa94DE7a4656D787667C749f7E1223D71E9FD88"),
			},
			Testnet: Contracts{
				Router:  common.HexToAddress("0x688d21b0B8Dc35971AF58cFF1F7Bf65639937860"),
				Factory: common.HexToAddress("0x2a496ec9e9bE22e66C61d4Eb9d316beaEE31A77b"),
			},
		},
	},
}
