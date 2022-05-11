package network

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDepositSend(t *testing.T) {
	t.Parallel()
	t.Run("pull image test", func(t *testing.T) {
		client, err := getDockerClient()
		require.NoError(t, err)
		ctx := context.Background()
		err = pullEtherealImage(ctx, client)
		require.NoError(t, err)
	})
	t.Run("send deposit test", func(t *testing.T) {
		client, err := getDockerClient()
		require.NoError(t, err)
		ctx := context.Background()
		nodeConf := DefaultL16NodeConfigs
		valSec := nodeConf.GetValSecrets()
		valSec.Eth1Data.WalletAddress = "0x7cBf71e554c72bdec8BF31d74Be3a9229C0CaF83"
		valSec.Eth1Data.WalletPrivKey = "f916f93513e14aa9cc9d2499515c28bbbcf788822c419869acdcc8923df13815"
		data, err := ParseDepositDataFromFile("../../assets/deposit_data.json")
		require.NoError(t, err)
		err = valSec.downloadEthereal(ctx, client)
		require.NoError(t, err)
		for _, x := range data {
			err = valSec.doDeposit(ctx, x, client)
			require.NoError(t, err)
		}
	})
}
