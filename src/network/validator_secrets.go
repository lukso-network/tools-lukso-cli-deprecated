package network

import (
	"github.com/lukso-network/lukso-cli/src/utils"
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
	ValidatorIndexFrom int64  `yaml:""`
	ValidatorIndexTo   int64  `yaml:""`
}

func (valSec *ValidatorCredentials) WriteToFile(fileName string) error {
	rawData, err := yaml.Marshal(valSec)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, rawData, os.ModePerm)
}

func (valSec *ValidatorCredentials) Print() {
	utils.ColoredPrintln("ValidatorMnemonic:", valSec.ValidatorMnemonic)
	utils.ColoredPrintln("Withdrawal Mnemonic:", valSec.WithdrawalMnemonic)
	utils.ColoredPrintln("Validator Index From:", valSec.ValidatorIndexFrom)
	utils.ColoredPrintln("Validator Index To:", valSec.ValidatorIndexTo)
}
