package beaconapi

import "fmt"

const (
	ParticipationPath = "/eth/v1alpha1/validators/participation?epoch=%d"
)

type ParticipationResponse struct {
	Epoch         string `json:"epoch"`
	Finalized     bool   `json:"finalized"`
	Participation struct {
		GlobalParticipationRate          float64 `json:"globalParticipationRate"`
		VotedEther                       string  `json:"votedEther"`
		EligibleEther                    string  `json:"eligibleEther"`
		CurrentEpochActiveGwei           string  `json:"currentEpochActiveGwei"`
		CurrentEpochAttestingGwei        string  `json:"currentEpochAttestingGwei"`
		CurrentEpochTargetAttestingGwei  string  `json:"currentEpochTargetAttestingGwei"`
		PreviousEpochActiveGwei          string  `json:"previousEpochActiveGwei"`
		PreviousEpochAttestingGwei       string  `json:"previousEpochAttestingGwei"`
		PreviousEpochTargetAttestingGwei string  `json:"previousEpochTargetAttestingGwei"`
		PreviousEpochHeadAttestingGwei   string  `json:"previousEpochHeadAttestingGwei"`
	} `json:"participation"`
}

func (c BeaconClient) Participation(epoch int64) (ParticipationResponse, error) {
	response := ParticipationResponse{}

	status, err := c.client.Get(&response, "", fmt.Sprintf(ParticipationPath, epoch))
	if err != nil {
		return response, fmt.Errorf("StatusCode: %d, Error: %s", status, err.Error())
	}

	return response, nil
}
