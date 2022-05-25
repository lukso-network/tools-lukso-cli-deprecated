package network

import "encoding/json"

type Bootnode struct {
	Consensus string `json:"consensus"`
	Execution string `json:"execution"`
}

type BootnodeUpdater struct {
	Chain Chain
}

func NewBootnodeUpdater(chain Chain) BootnodeUpdater {
	return BootnodeUpdater{chain}
}

func (b BootnodeUpdater) DownloadLatestBootnodes() ([]Bootnode, error) {
	url, err := getBootnodeUrl(b.Chain)
	if err != nil {
		return nil, err
	}
	bytes, err := downloadFileOverHttp(url.String())
	if err != nil {
		return nil, err
	}
	var result []Bootnode
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
