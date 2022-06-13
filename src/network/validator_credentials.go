package network

import (
	"github.com/lukso-network/lukso-cli/src/network/types"
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

func (c *ValidatorCredentials) WriteToFile(fileName string) error {
	rawData, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, rawData, os.ModePerm)
}

func (c *ValidatorCredentials) Print() {
	utils.ColoredPrintln("ValidatorMnemonic:", c.ValidatorMnemonic)
	utils.ColoredPrintln("Withdrawal Mnemonic:", c.WithdrawalMnemonic)
	utils.ColoredPrintln("Validator Index From:", c.ValidatorIndexFrom)
	utils.ColoredPrintln("Validator Index To:", c.ValidatorIndexTo)
}

func (c *ValidatorCredentials) CreateNodeRecovery() NodeRecovery {
	return NodeRecovery{
		ValidatorMnemonic:  c.ValidatorMnemonic,
		WithdrawalMnemonic: c.WithdrawalMnemonic,
		KeystoreIndexFrom:  c.ValidatorIndexFrom,
		KeystoreIndexTo:    c.ValidatorIndexTo,
	}
}

func (c *ValidatorCredentials) FromNodeRecovery(nr NodeRecovery) *ValidatorCredentials {
	c.ValidatorIndexTo = nr.KeystoreIndexTo
	c.ValidatorIndexFrom = nr.KeystoreIndexFrom
	c.WithdrawalMnemonic = nr.WithdrawalMnemonic
	c.ValidatorMnemonic = nr.ValidatorMnemonic
	return c
}

func (c *ValidatorCredentials) ValidatorRange() types.ValidatorRange {
	return types.ValidatorRange{
		From: c.ValidatorIndexFrom,
		To:   c.ValidatorIndexTo,
	}
}
