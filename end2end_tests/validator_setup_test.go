package end2end_tests

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestValidatorSetup(t *testing.T) {
	// TODO Skipping test for now
	t.Skip("skipping e2e tests for now")
	t.Parallel()
	t.Run("test downloader for validator tool", func(t *testing.T) {
		require.NoError(t, network.CheckAndDownloadValTool())
	})
	t.Run("test mnemonic generation", func(t *testing.T) {
		_, err := network.GetMnemonic()
		require.NoError(t, err)
	})
	t.Run("test secrets loader", func(t *testing.T) {
		expected := network.BetaDefaultValSecrets
		nodeconf := network.DefaultL16BetaNodeConfigs
		receivedFromFile := nodeconf.GetValSecrets()
		require.Equal(t, expected, receivedFromFile)
	})
	t.Run("genereate deposit data", func(t *testing.T) {
		mnemonic, err := network.GetMnemonic()
		require.NoError(t, err)
		valSec := network.BetaDefaultValSecrets
		valSec.ValidatorMnemonic = mnemonic
		valSec.WithdrawalMnemonic = mnemonic
		err = valSec.GenerateDepositData(5)
		require.NoError(t, err)
		err = os.RemoveAll(valSec.Deposit.DepositFileLocation)
		require.NoError(t, err)
	})
	t.Run("wallet creation test", func(t *testing.T) {
		viper.SetConfigFile("../../test_data/node_config.yaml")
		err := viper.ReadInConfig()
		require.NoError(t, err)
		mnemonic, err := network.GetMnemonic()
		require.NoError(t, err)
		valSec := network.BetaDefaultValSecrets
		valSec.ValidatorMnemonic = mnemonic
		err = valSec.GenerateWallet(5, "test1234")
		require.NoError(t, err)
		configs := network.MustGetNodeConfig()
		require.NoError(t, err)
		keyLocation, err := configs.GetKeyStorePath()
		require.NoError(t, err)
		err = os.RemoveAll(keyLocation)
		require.NoError(t, err)
	})
}
