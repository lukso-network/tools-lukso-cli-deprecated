package network

import (
	"github.com/lukso-network/lukso-cli/src"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatorTxStatus(t *testing.T) {
	t.Run("check validator txStatus", func(t *testing.T) {
		secrets, err := GetValSecrets(src.DefaultNetworkID)
		if err != nil {
			return
		}
		secrets.Deposit.DepositFileLocation = "../../assets/deposit_data.json"
		err = secrets.GetTxStatus()
		require.NoError(t, err)
	})
}
