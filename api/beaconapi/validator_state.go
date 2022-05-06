package beaconapi

import "fmt"

const (
	ValidatorStatePath = "/eth/v1/beacon/states/%s/validators/%s"
)

type ValidatorStateResponse struct {
	ExecutionOptimistic bool `json:"execution_optimistic"`
	Data                struct {
		Index     string `json:"index"`
		Balance   string `json:"balance"`
		Status    string `json:"status"`
		Validator struct {
			Pubkey                     string `json:"pubkey"`
			WithdrawalCredentials      string `json:"withdrawal_credentials"`
			EffectiveBalance           string `json:"effective_balance"`
			Slashed                    bool   `json:"slashed"`
			ActivationEligibilityEpoch string `json:"activation_eligibility_epoch"`
			ActivationEpoch            string `json:"activation_epoch"`
			ExitEpoch                  string `json:"exit_epoch"`
			WithdrawableEpoch          string `json:"withdrawable_epoch"`
		} `json:"validator"`
	} `json:"data"`
}

func (c BeaconClient) ValidatorState(pubKey string, epoch int64) (ValidatorStateResponse, error) {
	response := ValidatorStateResponse{}

	state := "head"
	if epoch > -1 {
		state = fmt.Sprintf("%d", epoch)
	}

	status, err := c.client.Get(&response, "", fmt.Sprintf(ValidatorStatePath, state, pubKey))
	if err != nil {
		return response, fmt.Errorf("StatusCode: %d, Error: %s", status, err.Error())
	}

	return response, nil
}
