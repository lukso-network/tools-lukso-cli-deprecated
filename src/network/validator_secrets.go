package network

import (
	"gopkg.in/yaml.v3"
	"os"
)

type DepositDetails struct {
	Amount              string `yaml:",omitempty"`
	ContractAddress     string `yaml:",omitempty"`
	DepositFileLocation string `yaml:",omitempty"`
	ForkVersion         string `yaml:",omitempty"`
}

type ValidatorCredentials struct {
	ValidatorMnemonic  string `yaml:",omitempty"`
	WithdrawalMnemonic string `yaml:",omitempty"`
}

type ValidatorSecretsV0 struct {
	ValidatorMnemonic  string `yaml:",omitempty"`
	WithdrawalMnemonic string `yaml:",omitempty"`
	ForkVersion        string `yaml:",omitempty"`

	Deposit *DepositDetails `yaml:",omitempty"`
}

func (valSec *ValidatorCredentials) WriteToFile(fileName string) error {
	rawData, err := yaml.Marshal(valSec)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, rawData, os.ModePerm)
}
