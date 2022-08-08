package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/scottrmalley/go-evm-constants/networks"
)

var SUSHISWAP Name = "sushiswap"

// sushiswapDex valus taken from the sushiswap sdk:
// https://github.com/sushiswap/sushiswap-sdk/blob/canary/src/constants/addresses.ts
var sushiswapDex = Dex{
	Name: SUSHISWAP,
	Networks: map[networks.Name]Contracts{
		networks.AVALANCHE_MAINNET: Contracts{
			Router:  common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"),
			Factory: common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		},
		networks.AVALANCHE_FUJI: Contracts{
			Router:  common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"),
			Factory: common.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
		},
		networks.BSC_MAINNET: Contracts{
			Router:  common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"),
			Factory: common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		},
		networks.BSC_TESTNET: Contracts{
			Router:  common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"),
			Factory: common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		},
		networks.POLYGON_MAINNET: Contracts{
			Router:  common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"),
			Factory: common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		},
		networks.POLYGON_TESTNET: Contracts{
			Router:  common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"),
			Factory: common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		},
		networks.ETHEREUM_MAINNET: Contracts{
			Router:  common.HexToAddress("0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"),
			Factory: common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"),
		},
		networks.ETHEREUM_GOERLI: Contracts{
			Router:  common.HexToAddress("0x1b02dA8Cb0d097eB8D57A175b88c7D8b47997506"),
			Factory: common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4"),
		},
	},
}
