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

	config, err := GetLoadedNodeConfigs()
	require.NoError(t, err)

	t.Run("retrieve node name without any argument", func(t *testing.T) {
		nodeName := config.getNodeName("")
		if config.Node == nil || config.Node.Name == "" {
			hostName, _ := os.Hostname()
			require.Equal(t, hostName, nodeName)
		} else {
			require.Equal(t, config.Node.Name, nodeName)
		}
	})
	t.Run("retrieve node name with argument", func(t *testing.T) {
		require.Equal(t, "myNode", config.getNodeName("myNode"))
	})

	t.Run("retrieve config directory", func(t *testing.T) {
		if config.Configs != nil && config.Configs.Volume != "" {
			configPath, err := config.getConfigPath()
			require.NoError(t, err)
			require.Equal(t, config.Configs.Volume, configPath)
		}
		if config.Configs != nil {
			tempVol := config.getConfigs().getVolume()
			config.Configs.Volume = ""
			homeDir, err := os.UserHomeDir()
			require.NoError(t, err)
			luksoConfigPath := path.Join(homeDir, ".lukso_configs")
			configPath, err := config.getConfigPath()
			require.NoError(t, err)
			require.Equal(t, luksoConfigPath, configPath)
			config.Configs.Volume = tempVol
		}
		tempConf := config.getConfigs()
		config.Configs = nil
		homeDir, err := os.UserHomeDir()
		require.NoError(t, err)
		luksoConfigPath := path.Join(homeDir, ".lukso_configs")
		configPath, err := config.getConfigPath()
		require.NoError(t, err)
		require.Equal(t, luksoConfigPath, configPath)
		config.Configs = tempConf
	})
	t.Run("retrieve execution data", func(t *testing.T) {
		expectedData := config.Execution
		require.Equal(t, expectedData, config.getExecutionData())
	})
	t.Run("retrieve consensus data", func(t *testing.T) {
		expectedData := config.Consensus
		require.Equal(t, expectedData, config.getConsensusData())
	})
	t.Run("retrieve validator data", func(t *testing.T) {
		expectedData := config.Validator
		require.Equal(t, expectedData, config.getValidatorData())
	})
	t.Run("retrieve stat address", func(t *testing.T) {
		require.Equal(t, config.Execution.StatsAddress, getDataFromContainer(config.getExecutionData(), keyStatsAddress))
		require.Equal(t, config.Consensus.StatsAddress, getDataFromContainer(config.getConsensusData(), keyStatsAddress))
	})
	t.Run("retrieve geth ports", func(t *testing.T) {
		_, err := config.getGethHttpPort()
		require.NoError(t, err)
		_, err = config.gethGethPeerPort()
		require.NoError(t, err)
	})
}
