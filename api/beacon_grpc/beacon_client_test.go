package beacon_grpc

import (
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBeaconClient_GetValidatorStatus(t *testing.T) {
	t.Skip("skip e2e tests")
	nodeconf := network.DefaultL16BetaNodeConfigs
	valSec := nodeconf.GetCredentials()
	valClient, err := NewBeaconClient(valSec.Eth2Data.GRPCEndPoint)
	require.NoError(t, err)
	depositData, err := network.ParseDepositDataFromFile("../../test_data/deposit_data.json")
	require.NoError(t, err)
	for _, depData := range depositData {
		_, err := valClient.GetValidatorStatus([]byte(depData.PubKey))
		require.NoError(t, err)
	}
}
