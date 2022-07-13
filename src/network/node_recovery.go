package network

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type NodeRecovery struct {
	ValidatorCredentials struct {
		ValidatorMnemonic  string `json:"validatorMnemonic"`
		WithdrawalMnemonic string `json:"withdrawalMnemonic"`
		KeystoreIndexFrom  int64  `json:"keystoreIndexFrom"`
		KeystoreIndexTo    int64  `json:"keystoreIndexTo"`
	} `json:"validatorCredentials"`
	TransactionWallet struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	}
}

func LoadNodeRecovery(source string) (*NodeRecovery, error) {
	bytes, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}
	node := &NodeRecovery{}
	err = json.Unmarshal(bytes, node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (nr NodeRecovery) Save() error {
	bytes, err := json.Marshal(nr)
	if err != nil {
		return err
	}
	return os.WriteFile(NodeRecoveryFileLocation, bytes, os.ModePerm)
}

func (nr NodeRecovery) Append() error {
	bytes, err := json.Marshal(nr)
	if err != nil {
		return err
	}
	return os.WriteFile(NodeRecoveryFileLocation, bytes, os.ModeAppend)
}
