package end2end_tests

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatorTxStatus(t *testing.T) {
	// TODO Skipping test for now
	t.Skip("skipping e2e tests for now")
	t.Run("check validator txStatus", func(t *testing.T) {
		nodeConf := network.DefaultL16BetaNodeConfigs
		secrets := nodeConf.GetValSecrets()
		secrets.Deposit.DepositFileLocation = "../../test_data/deposit_data.json"
		err := secrets.GetTxStatus()
		require.NoError(t, err)
	})
}
