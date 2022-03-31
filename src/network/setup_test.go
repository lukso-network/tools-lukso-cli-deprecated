package network

import (
	"github.com/lukso-network/lukso-cli/src"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDownloadClient(t *testing.T) {
	t.Run("download makefile", func(t *testing.T) {
		viper.Set(src.ViperKeyNetworkName, src.DefaultNetworkID)
		require.NoError(t, downloadNetworkSetupFiles())
	})
	t.Run("modify env file", func(t *testing.T) {
		require.NoError(t, modifyEnvFile("test-host"))
	})
	t.Run("download config files", func(t *testing.T) {
		viper.Set(src.ViperKeyNetworkName, src.DefaultNetworkID)
		require.NoError(t, downloadConfigFiles())
	})
}
