package networks

import (
	"errors"
	"math/big"
)

var ErrNetworkNotFound = errors.New("could not find network")
var ErrChainIdNotFound = errors.New("could not find network with specified chain id")
var ErrCantRegisterEmptyNetwork = errors.New("can't register empty network")

type Name string

type Network struct {
	RpcUrl  string
	ChainId *big.Int
	Name    Name
	Testnet bool
}

var networkMap = map[Name]Network{
	ETHEREUM_MAINNET:  ethereum,
	ETHEREUM_GOERLI:   ethereumGoerli,
	AVALANCHE_MAINNET: avalancheMainnet,
	AVALANCHE_FUJI:    avalancheFuji,
	BSC_MAINNET:       bscMainnet,
	BSC_TESTNET:       bscTestnet,
	POLYGON_MAINNET:   polygonMainnet,
	POLYGON_TESTNET:   polygonTestnet,
}

func GetNetwork(name Name) (Network, error) {
	if net, ok := networkMap[name]; ok {
		return net, nil
	} else {
		return Network{}, ErrNetworkNotFound
	}
}

func RegisterNetwork(network Network) error {
	if network.Name == "" {
		return ErrCantRegisterEmptyNetwork
	}
	networkMap[network.Name] = network
	return nil
}

func NameFrom(chainId *big.Int) (Name, error) {
	for name, network := range networkMap {
		if network.ChainId.String() == chainId.String() {
			return name, nil
		}
	}
	return "", ErrChainIdNotFound
}

func ChainIdFrom(name Name) (*big.Int, error) {
	if network, ok := networkMap[name]; !ok {
		return nil, ErrNetworkNotFound
	} else {
		return network.ChainId, nil
	}
}

func SupportedNetworks() []Name {
	var supported []Name
	for name := range networkMap {
		supported = append(supported, name)
	}
	return supported
}

func Supported(name string) (Name, bool) {
	networkName := Name(name)
	if _, ok := networkMap[networkName]; ok {
		return networkName, true
	}
	return "", false
}
