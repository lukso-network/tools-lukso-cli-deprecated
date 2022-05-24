package network

import (
	"gopkg.in/yaml.v3"
	"os"
)

type DepositDetails struct {
	Amount              string `yaml:",omitempty"`
	ContractAddress     string `yaml:",omitempty"`
	DepositFileLocation string `yaml:",omitempty"`
	Force               bool   `yaml:",omitempty"`
}

type Eth1Details struct {
	WalletAddress string `yaml:",omitempty"`
	WalletPrivKey string `yaml:",omitempty"`
}

type Eth2Details struct {
	GRPCEndPoint string `yaml:",omitempty"`
}

type ValidatorSecrets struct {
	ValidatorMnemonic  string `yaml:",omitempty"`
	WithdrawalMnemonic string `yaml:",omitempty"`
	ForkVersion        string `yaml:",omitempty"`

	Deposit  *DepositDetails `yaml:",omitempty"`
	Eth1Data *Eth1Details    `yaml:",omitempty"`
	Eth2Data *Eth2Details    `yaml:",omitempty"`
}

func (valSec *ValidatorSecrets) WriteToFile(fileName string) error {
	rawData, err := yaml.Marshal(valSec)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, rawData, os.ModePerm)
}
