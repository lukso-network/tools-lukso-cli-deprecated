package network

import (
	"github.com/lukso-network/lukso-cli/src"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestDownloadClient(t *testing.T) {
	t.Run("modify env file", func(t *testing.T) {
		viper.SetConfigFile("../../node_config.yaml")
		err := viper.ReadInConfig()
		require.NoError(t, err)

		require.NoError(t, GenerateEnvFile("test-host"))
	})
	t.Run("generate node config", func(t *testing.T) {
		err := GenerateDefaultNodeConfigs(src.DefaultNetworkID)
		require.NoError(t, err)

		luksoConfigHomePath := "./node_config.yaml"

		var actualConfig NodeConfigs
		rawData, err := os.ReadFile(luksoConfigHomePath)
		require.NoError(t, err)

		err = yaml.Unmarshal(rawData, &actualConfig)
		require.NoError(t, err)
		require.Equal(t, DefaultL16NodeConfigs, &actualConfig)
	})
}
