package network

import (
	"encoding/json"
	"os"
)

type DepositData struct {
	Account               string `json:"account"`
	DepositDataRoot       string `json:"deposit_data_root"`
	PubKey                string `json:"pubkey"`
	Signature             string `json:"signature"`
	Value                 int    `json:"value"`
	Version               int    `json:"version"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
}

func ParseDepositDataFromFile(fileName string) ([]*DepositData, error) {
	rawData, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var parsedData []*DepositData
	err = json.Unmarshal(rawData, &parsedData)
	if err != nil {
		return nil, err
	}
	return parsedData, nil
}
