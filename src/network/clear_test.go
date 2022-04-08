package network

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClear(t *testing.T) {
	viper.SetConfigFile("../../node_config.yaml")
	err := viper.ReadInConfig()
	require.NoError(t, err)
	t.Run("test clear data (where only one file existed)", func(t *testing.T) {
		valDataPath := getValidatorDataVolume()
		err = os.MkdirAll(valDataPath, os.ModePerm)
		require.NoError(t, err)

		err = Clear()
		require.NoError(t, err)
		require.NoFileExists(t, valDataPath)
	})
	t.Run("test clear data (where all files existed)", func(t *testing.T) {
		valDataPath := getValidatorDataVolume()
		err = os.MkdirAll(valDataPath, os.ModePerm)
		require.NoError(t, err)

		consensusDataPath := getConsensusDataVolume()
		err = os.MkdirAll(consensusDataPath, os.ModePerm)
		require.NoError(t, err)

		executionDataPath := getExecutionDataVolume()
		err = os.MkdirAll(executionDataPath, os.ModePerm)
		require.NoError(t, err)

		err = Clear()
		require.NoError(t, err)
		require.NoFileExists(t, valDataPath)
		require.NoFileExists(t, consensusDataPath)
		require.NoFileExists(t, executionDataPath)
	})
}
