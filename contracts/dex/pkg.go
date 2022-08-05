package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scottrmalley/go-evm-constants/networks"
)

var (
	ErrDexNotFound         = errors.New("no dex with provided name")
	ErrNetworkNotSupported = errors.New("the given dex is not supported on the network")
)

type Name string

type Contracts struct {
	Router  common.Address
	Factory common.Address
}

type NetworkDex struct {
	Mainnet Contracts
	Testnet Contracts
}

type Dex struct {
	Name     Name
	Networks map[networks.Name]NetworkDex
}

var dexMap = map[Name]Dex{}

func GetContracts(name Name, network networks.Name, testnet bool) (Contracts, error) {
	dex, ok := dexMap[name]
	if !ok {
		return Contracts{}, ErrDexNotFound
	}
	contracts, ok := dex.Networks[network]
	if !ok {
		return Contracts{}, ErrNetworkNotSupported
	}
	if testnet && contracts.Testnet != (Contracts{}) {
		return contracts.Testnet, nil
	} else if contracts.Mainnet != (Contracts{}) {
		return contracts.Mainnet, nil
	}
	return Contracts{}, ErrNetworkNotSupported
}
