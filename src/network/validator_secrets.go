package network

import (
	"gopkg.in/yaml.v3"
	"os"
)

type DepositDetails struct {
	Amount              string `yaml:""`
	ContractAddress     string `yaml:""`
	DepositFileLocation string `yaml:""`
	ForkVersion         string `yaml:""`
}

type ValidatorCredentials struct {
	ValidatorMnemonic  string `yaml:""`
	WithdrawalMnemonic string `yaml:""`
}

type ValidatorSecretsV0 struct {
	ValidatorMnemonic  string `yaml:""`
	WithdrawalMnemonic string `yaml:""`
	ForkVersion        string `yaml:""`

	Deposit *DepositDetails `yaml:""`
}

func (valSec *ValidatorCredentials) WriteToFile(fileName string) error {
	rawData, err := yaml.Marshal(valSec)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, rawData, os.ModePerm)
}
