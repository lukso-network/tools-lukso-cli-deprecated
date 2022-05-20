package end2end_tests

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClear(t *testing.T) {
	// TODO Skipping test for now
	t.Skip("skipping e2e tests for now")
	viper.SetConfigFile("../../node_config.yaml")
	err := viper.ReadInConfig()
	require.NoError(t, err)
	t.Run("test clear data (where only one file existed)", func(t *testing.T) {
		valDataPath, err := network.GetValidatorDataVolume()
		require.NoError(t, err)
		err = os.MkdirAll(valDataPath, os.ModePerm)
		require.NoError(t, err)

		err = network.Clear()
		require.NoError(t, err)
		require.NoFileExists(t, valDataPath)
	})
	t.Run("test clear data (where all files existed)", func(t *testing.T) {
		valDataPath, err := network.GetValidatorDataVolume()
		require.NoError(t, err)

		err = os.MkdirAll(valDataPath, os.ModePerm)
		require.NoError(t, err)

		consensusDataPath, err := network.GetConsensusDataVolume()
		require.NoError(t, err)

		err = os.MkdirAll(consensusDataPath, os.ModePerm)
		require.NoError(t, err)

		executionDataPath, err := network.GetExecutionDataVolume()
		require.NoError(t, err)

		err = os.MkdirAll(executionDataPath, os.ModePerm)
		require.NoError(t, err)

		err = network.Clear()
		require.NoError(t, err)
		require.NoFileExists(t, valDataPath)
		require.NoFileExists(t, consensusDataPath)
		require.NoFileExists(t, executionDataPath)
	})
}
