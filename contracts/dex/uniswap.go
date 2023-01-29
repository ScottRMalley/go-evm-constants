package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/scottrmalley/go-evm-constants/networks"
)

var UNISWAP Name = "uniswap"

// uniswapDex values pulled from the usiswap docs:
// https://docs.uniswap.org/contracts/v3/reference/deployments
var uniswapDex = Dex{
	Name: UNISWAP,
	Networks: map[networks.Name]Contracts{
		networks.ETHEREUM_MAINNET: Contracts{
			Router:  common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564"),
			Factory: common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
		},
		networks.ETHEREUM_GOERLI: Contracts{
			Router:  common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564"),
			Factory: common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
		},
		networks.POLYGON_MAINNET: Contracts{
			Router:  common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564"),
			Factory: common.HexToAddress("0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32"),
		},
		networks.POLYGON_TESTNET: Contracts{
			Router:  common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564"),
			Factory: common.HexToAddress("0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32"),
		},
	},
}
