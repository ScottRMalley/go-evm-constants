package networks

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"math/big"
	"testing"
)

type NetworkTestSuite struct {
	suite.Suite
}

func TestNetworkTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkTestSuite))
}

func (s *NetworkTestSuite) TestNetworkNameDefaults() {
	t := s.T()
	mainnetNames := []Name{BSC_MAINNET, ETHEREUM_MAINNET, AVALANCHE_MAINNET, POLYGON_MAINNET}
	testnetNames := []Name{BSC_TESTNET, ETHEREUM_GOERLI, AVALANCHE_FUJI, POLYGON_TESTNET}

	t.Run("should have properly configured mainnets", func(t *testing.T) {
		for _, name := range mainnetNames {
			mainnet, err := GetNetwork(name)
			require.NoError(t, err, "should not produce an error when fetching known network")
			assert.NotEmpty(t, mainnet, "should not return empty network")
			assert.False(t, mainnet.Testnet, "mainnet should not be marked as testnet")
		}
	})

	t.Run("should have properly configured testnets", func(t *testing.T) {
		for _, name := range testnetNames {
			testnet, err := GetNetwork(name)
			require.NoError(t, err, "should not return an error when fetching known network")
			assert.NotEmpty(t, testnet, "should not return an empty network")
			assert.Truef(t, testnet.Testnet, "testnet should be marked as testnet: %s", name)
		}
	})
}

func (s *NetworkTestSuite) TestCustomNetworkNames() {
	t := s.T()
	customNetworkName := Name("custom")

	t.Run("should fail to register empty network", func(t *testing.T) {
		err := RegisterNetwork(Network{})
		require.Error(t, err, "should return error on failed network registration")
	})

	t.Run("should register a properly configured blockchain", func(t *testing.T) {
		network := Network{
			RpcUrl:  "http://rpc.custom.io/",
			ChainId: big.NewInt(1234),
			Name:    customNetworkName,
			Testnet: false,
		}

		err := RegisterNetwork(network)
		require.NoError(t, err)

		custom, err := GetNetwork(customNetworkName)
		require.NoError(t, err, "should not return an err for registered network")
		require.NotEmpty(t, custom, "should not return empty blockchain for registered network")

		assert.Equal(t, custom.Name, customNetworkName)
	})
}
