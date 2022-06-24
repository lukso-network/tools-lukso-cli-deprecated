package beaconapi

import "fmt"

const (
	IdentityPath = "eth/v1/node/identity"
)

type IdentityResponse struct {
	Data struct {
		PeerID             string   `json:"peer_id"`
		Enr                string   `json:"enr"`
		P2PAddresses       []string `json:"p2p_addresses"`
		DiscoveryAddresses []string `json:"discovery_addresses"`
		Metadata           struct {
			SeqNumber string `json:"seq_number"`
			Attnets   string `json:"attnets"`
			Syncnets  string `json:"syncnets"`
		} `json:"metadata"`
	} `json:"data"`
}

func (c BeaconClient) Identity() (IdentityResponse, error) {
	response := IdentityResponse{}

	status, err := c.client.Get(&response, "", IdentityPath)
	if err != nil {
		return response, fmt.Errorf("StatusCode: %d, Error: %s", status, err.Error())
	}

	return response, nil
}
