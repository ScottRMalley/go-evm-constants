package dex

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/scottrmalley/go-evm-constants/networks"
)

var (
	ErrDexNotFound         = errors.New("no dex with provided name")
	ErrNetworkNotSupported = errors.New("the given dex is not supported on the network")
	ErrEmptyDex            = errors.New("cannot register empty dex")
)

type Name string

type Contracts struct {
	Router  common.Address
	Factory common.Address
}

type Dex struct {
	Name     Name
	Networks map[networks.Name]Contracts
}

func (d Dex) Supports(network networks.Name) bool {
	_, ok := d.Networks[network]
	return ok
}

var dexMap = map[Name]Dex{
	PANGOLIN:   pangolinDex,
	SUSHISWAP:  sushiswapDex,
	TRADER_JOE: traderJoeDex,
	UNISWAP:    uniswapDex,
}

func GetContracts(name Name, network networks.Name) (Contracts, error) {
	dex, ok := dexMap[name]
	if !ok {
		return Contracts{}, ErrDexNotFound
	}
	contracts, ok := dex.Networks[network]
	if !ok {
		return Contracts{}, ErrNetworkNotSupported
	}
	return contracts, nil
}

func RegisterDex(dex Dex) error {
	if dex.Name == "" {
		return ErrEmptyDex
	}
	dexMap[dex.Name] = dex
	return nil
}

func SupportsNetwork(name Name, network networks.Name) (bool, error) {
	dex, ok := dexMap[name]
	if !ok {
		return false, ErrDexNotFound
	}
	return dex.Supports(network), nil
}
