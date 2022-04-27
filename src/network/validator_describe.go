package network

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/utils/beacon_grpc"
)

func (valSec *ValidatorSecrets) GetTxStatus() error {
	endPoint := valSec.Eth2Data.GRPCEndPoint
	validatorClient, err := beacon_grpc.NewBeaconClient(endPoint)
	if err != nil {
		return err
	}
	depositData, err := ParseDepositDataFromFile(valSec.Deposit.DepositFileLocation)
	if err != nil {
		return err
	}
	for _, depData := range depositData {
		status, err := validatorClient.GetValidatorStatus([]byte(depData.PubKey))
		if err != nil {
			return err
		}
		fmt.Println("Public Key:", depData.PubKey, "Status:", status)
	}
	return nil
}
