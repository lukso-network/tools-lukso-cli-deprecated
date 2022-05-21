package beaconapi

import "fmt"

const (
	ChainHeadPath = "eth/v1alpha1/beacon/chainhead"
)

type ChainHeadResponse struct {
	HeadSlot                   string `json:"headSlot"`
	HeadEpoch                  string `json:"headEpoch"`
	HeadBlockRoot              string `json:"headBlockRoot"`
	FinalizedSlot              string `json:"finalizedSlot"`
	FinalizedEpoch             string `json:"finalizedEpoch"`
	FinalizedBlockRoot         string `json:"finalizedBlockRoot"`
	JustifiedSlot              string `json:"justifiedSlot"`
	JustifiedEpoch             string `json:"justifiedEpoch"`
	JustifiedBlockRoot         string `json:"justifiedBlockRoot"`
	PreviousJustifiedSlot      string `json:"previousJustifiedSlot"`
	PreviousJustifiedEpoch     string `json:"previousJustifiedEpoch"`
	PreviousJustifiedBlockRoot string `json:"previousJustifiedBlockRoot"`
}

func (c BeaconClient) ChainHead() (ChainHeadResponse, error) {
	response := ChainHeadResponse{}

	status, err := c.client.Get(&response, "", ChainHeadPath)
	if err != nil {
		return response, fmt.Errorf("StatusCode: %d, Error: %s", status, err.Error())
	}

	return response, nil
}
