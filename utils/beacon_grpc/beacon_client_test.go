package beacon_grpc

import (
	"github.com/lukso-network/lukso-cli/src"
	"github.com/lukso-network/lukso-cli/src/network"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBeaconClient_GetValidatorStatus(t *testing.T) {
	valSec, err := network.GetValSecrets(src.DefaultNetworkID)
	require.NoError(t, err)
	valClient, err := NewBeaconClient(valSec.Eth2Data.GRPCEndPoint)
	require.NoError(t, err)
	depositData, err := network.ParseDepositDataFromFile("../../assets/deposit_data.json")
	require.NoError(t, err)
	for _, depData := range depositData {
		_, err := valClient.GetValidatorStatus([]byte(depData.PubKey))
		require.NoError(t, err)
	}
}
