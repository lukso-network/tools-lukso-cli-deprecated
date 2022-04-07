package network

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestParser(t *testing.T) {
	t.Parallel()
	viper.SetConfigFile("../../node_config.yaml")
	require.NoError(t, viper.ReadInConfig())
	t.Run("retrieve node name without any argument", func(t *testing.T) {
		nodeName := getNodeName("")
		if getNodeInfo()[keyName] == nil {
			hostName, _ := os.Hostname()
			require.Equal(t, hostName, nodeName)
		} else {
			require.Equal(t, getNodeInfo()[keyName].(string), nodeName)
		}
	})
	t.Run("retrieve node name with argument", func(t *testing.T) {
		require.Equal(t, "myNode", getNodeName("myNode"))
	})

	t.Run("retrieve config directory", func(t *testing.T) {
		tempContainer := viper.Get(keyConfigs)
		viper.Set(keyConfigs, map[string]interface{}{})
		homeDir, err := os.UserHomeDir()
		require.NoError(t, err)
		expectedPath := path.Join(homeDir, ".lukso_config")
		returnedPath, err := getConfigPath()
		require.Equal(t, expectedPath, returnedPath)
		viper.Set(keyConfigs, tempContainer)
	})
	t.Run("retrieve execution data", func(t *testing.T) {
		expectedData := viper.GetStringMapString(keyExecution)
		require.Equal(t, expectedData, getExecutionData())
	})
	t.Run("retrieve consensus data", func(t *testing.T) {
		expectedData := viper.GetStringMapString(keyConsensus)
		require.Equal(t, expectedData, getConsensusData())
	})
	t.Run("retrieve validator data", func(t *testing.T) {
		expectedData := viper.GetStringMapString(keyValidator)
		require.Equal(t, expectedData, getValidatorData())
	})
	t.Run("retrieve stat address", func(t *testing.T) {
		if dataContainer := viper.GetStringMapString(keyExecution); dataContainer != nil {
			expectedData := dataContainer[keyStatsAddress]
			require.Equal(t, expectedData, getDataFromContainer(dataContainer, keyStatsAddress))
		}
		if dataContainer := viper.GetStringMapString(keyConsensus); dataContainer != nil {
			expectedData := dataContainer[keyStatsAddress]
			require.Equal(t, expectedData, getDataFromContainer(dataContainer, keyStatsAddress))
		}
	})
	t.Run("retrieve geth ports", func(t *testing.T) {
		_, err := getGethHttpPort()
		require.NoError(t, err)
		_, err = gethGethPeerPort()
		require.NoError(t, err)
	})
}
