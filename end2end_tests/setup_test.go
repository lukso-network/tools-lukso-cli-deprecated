package end2end_tests

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestDownloadClient(t *testing.T) {
	// TODO Skipping test for now
	t.Skip("skipping e2e tests for now")
	t.Run("modify env file", func(t *testing.T) {
		viper.SetConfigFile("../../test_data/node_config.yaml")
		err := viper.ReadInConfig()
		require.NoError(t, err)

		require.NoError(t, network.GenerateEnvFile("test-host"))
	})
	t.Run("generate node config", func(t *testing.T) {
		err := network.GenerateDefaultNodeConfigs(network.L16Beta)
		require.NoError(t, err)

		luksoConfigHomePath := "./node_config.yaml"

		var actualConfig network.NodeConfigs
		rawData, err := os.ReadFile(luksoConfigHomePath)
		require.NoError(t, err)

		err = yaml.Unmarshal(rawData, &actualConfig)
		require.NoError(t, err)
		require.Equal(t, network.DefaultL16BetaNodeConfigs, &actualConfig)
	})
}
