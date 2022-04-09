package network

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"testing"
)

func TestDownloadClient(t *testing.T) {
	t.Run("download makefile", func(t *testing.T) {
		viper.Set(src.ViperKeyNetworkName, src.DefaultNetworkID)
		require.NoError(t, downloadNetworkSetupFiles())
	})
	t.Run("modify env file", func(t *testing.T) {
		viper.SetConfigFile("../../node_config.yaml")
		err := viper.ReadInConfig()
		require.NoError(t, err)

		require.NoError(t, generateEnvFile("test-host"))
	})
	t.Run("download config files", func(t *testing.T) {
		viper.Set(src.ViperKeyNetworkName, src.DefaultNetworkID)
		require.NoError(t, downloadConfigFiles())
	})
	t.Run("generate node config", func(t *testing.T) {
		err := GenerateDefaultNodeConfigs(src.DefaultNetworkID)
		require.NoError(t, err)

		userHomeDir, err := os.UserHomeDir()
		require.NoError(t, err)
		luksoConfigHomePath := path.Join(userHomeDir, fmt.Sprintf(".lukso_%s", src.DefaultNetworkID), "node_configs.yaml")

		var actualConfig NodeConfigs
		rawData, err := os.ReadFile(luksoConfigHomePath)
		require.NoError(t, err)

		err = yaml.Unmarshal(rawData, &actualConfig)
		require.NoError(t, err)
		require.Equal(t, DefaultL16NodeConfigs, &actualConfig)
	})
}
