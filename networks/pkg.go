package networks

import (
	"errors"
	"math/big"
)

var ErrNetworkNotFound = errors.New("could not find network")
var ErrCantRegisterEmptyBlockchain = errors.New("can't register empty blockchain")

type Name string

type Network struct {
	RpcUrl  string
	ChainId *big.Int
	Name    Name
	Testnet bool
}

type Blockchain struct {
	Mainnet Network
	Testnet Network
}

var blockchainMap = map[Name]Blockchain{
	ETHEREUM:  ETHEREUM_BLOCKCHAIN,
	AVALANCHE: AVALANCHE_BLOCKCHAIN,
	BSC:       BSC_BLOCKCHAIN,
	POLYGON:   POLYGON_BLOCKCHAIN,
}

func GetNetwork(name Name, testnet bool) (Network, error) {
	if bc, ok := blockchainMap[name]; !ok {
		return Network{}, ErrNetworkNotFound
	} else {
		if testnet && bc.Testnet != (Network{}) {
			return bc.Testnet, nil
		} else if bc.Mainnet != (Network{}) {
			return bc.Mainnet, nil
		} else {
			return Network{}, ErrNetworkNotFound
		}
	}
}

func RegisterBlockchain(name Name, bc Blockchain) error {
	if bc.Mainnet == (Network{}) && bc.Testnet == (Network{}) {
		return ErrCantRegisterEmptyBlockchain
	}
	blockchainMap[name] = bc
	return nil
}
