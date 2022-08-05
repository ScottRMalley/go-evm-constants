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
	names := []Name{BSC, ETHEREUM, AVALANCHE, POLYGON}

	t.Run("should have properly configured mainnets", func(t *testing.T) {
		for _, name := range names {
			mainnet, err := GetNetwork(name, false)
			require.NoError(t, err, "should not produce an error when fetching known network")
			assert.NotEmpty(t, mainnet, "should not return empty network")
			assert.False(t, mainnet.Testnet, "mainnet should not be marked as testnet")
		}
	})

	t.Run("should have properly configured testnets", func(t *testing.T) {
		for _, name := range names {
			testnet, err := GetNetwork(name, true)
			require.NoError(t, err, "should not return an error when fetching known network")
			assert.NotEmpty(t, testnet, "should not return an empty network")
			assert.Truef(t, testnet.Testnet, "testnet should be marked as testnet: %s", name)
		}
	})
}

func (s *NetworkTestSuite) TestCustomNetworkNames() {
	t := s.T()
	customNetworkName := Name("custom")

	t.Run("should fail to register empty blockchain", func(t *testing.T) {
		err := RegisterBlockchain(customNetworkName, Blockchain{})
		require.Error(t, err, "should return error on failed network registration")

		custom, err := GetNetwork(customNetworkName, true)
		require.Error(t, err, "should return err for unregistered network")
		assert.Empty(t, custom, "should return empty for unknown network")
	})

	t.Run("should register a properly configured blockchain", func(t *testing.T) {
		bc := Blockchain{
			Mainnet: Network{
				RpcUrl:  "http://rpc.custom.io/",
				ChainId: big.NewInt(1234),
				Name:    customNetworkName,
				Testnet: false,
			},
		}
		err := RegisterBlockchain(customNetworkName, bc)
		require.NoError(t, err)

		custom, err := GetNetwork(customNetworkName, false)
		require.NoError(t, err, "should not return an err for registered network")
		require.NotEmpty(t, custom, "should not return empty blockchain for registered network")

		assert.Equal(t, custom.Name, customNetworkName)
	})
}
