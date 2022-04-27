package network

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestValidatorSetup(t *testing.T) {
	t.Parallel()
	t.Run("test downloader for validator tool", func(t *testing.T) {
		require.NoError(t, checkAndDownloadValTool())
	})
	t.Run("test mnemonic generation", func(t *testing.T) {
		_, err := getMnemonic()
		require.NoError(t, err)
	})
	t.Run("test secrets loader", func(t *testing.T) {
		expected := BetaDefaultValSecrets
		require.NoError(t, expected.WriteToFile("./.secrets.yaml"))
		receivedFromFile, err := GetValSecretsFromFile("./.secrets.yaml")
		require.NoError(t, err)
		require.Equal(t, expected, receivedFromFile)
		err = os.RemoveAll("./.secrets.yaml")
		require.NoError(t, err)
	})
	t.Run("genereate deposit data", func(t *testing.T) {
		mnemonic, err := getMnemonic()
		require.NoError(t, err)
		valSec := BetaDefaultValSecrets
		valSec.ValidatorMnemonic = mnemonic
		valSec.WithdrawalMnemonic = mnemonic
		err = valSec.GenerateDepositData(5)
		require.NoError(t, err)
		err = os.RemoveAll(valSec.Deposit.DepositFileLocation)
		require.NoError(t, err)
	})
	t.Run("wallet creation test", func(t *testing.T) {
		viper.SetConfigFile("../../node_config.yaml")
		err := viper.ReadInConfig()
		require.NoError(t, err)
		mnemonic, err := getMnemonic()
		require.NoError(t, err)
		valSec := BetaDefaultValSecrets
		valSec.ValidatorMnemonic = mnemonic
		err = valSec.GenerateWallet(5, "test1234")
		require.NoError(t, err)
		configs, err := GetLoadedNodeConfigs()
		require.NoError(t, err)
		keyLocation, err := configs.getKeyStorePath()
		require.NoError(t, err)
		err = os.RemoveAll(keyLocation)
		require.NoError(t, err)
	})
}
