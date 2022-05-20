package network

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatorTxStatus(t *testing.T) {
	t.Run("check validator txStatus", func(t *testing.T) {
		nodeConf := DefaultL16BetaNodeConfigs
		secrets := nodeConf.GetValSecrets()
		secrets.Deposit.DepositFileLocation = "../../assets/deposit_data.json"
		err := secrets.GetTxStatus()
		require.NoError(t, err)
	})
}
