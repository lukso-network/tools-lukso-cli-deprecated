package end2end_tests

import (
	"context"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDepositSend(t *testing.T) {
	// TODO Skipping test for now
	t.Skip("skipping e2e tests for now")
	t.Parallel()
	t.Run("pull image test", func(t *testing.T) {
		client, err := network.GetDockerClient()
		require.NoError(t, err)
		ctx := context.Background()
		err = network.PullEtherealImage(ctx, client)
		require.NoError(t, err)
	})
	t.Run("send deposit test", func(t *testing.T) {
		client, err := network.GetDockerClient()
		require.NoError(t, err)
		ctx := context.Background()
		nodeConf := network.DefaultL16BetaNodeConfigs
		valSec := nodeConf.GetValSecrets()
		valSec.Eth1Data.WalletAddress = "0x7cBf71e554c72bdec8BF31d74Be3a9229C0CaF83"
		valSec.Eth1Data.WalletPrivKey = "f916f93513e14aa9cc9d2499515c28bbbcf788822c419869acdcc8923df13815"
		data, err := network.ParseDepositDataFromFile("../../assets/deposit_data.json")
		require.NoError(t, err)
		err = valSec.DownloadEthereal(ctx, client)
		require.NoError(t, err)
		for _, x := range data {
			err = valSec.DoDeposit(ctx, x, client, "")
			require.NoError(t, err)
		}
	})
}
