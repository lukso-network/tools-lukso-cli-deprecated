package network

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/network/types"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"path"
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

func (c *ValidatorCredentials) IsEmpty() bool {
	return c.ValidatorMnemonic == "" || c.WithdrawalMnemonic == ""
}

func (c *ValidatorCredentials) GenerateMnemonic() error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}
	useExisting, err := UseExistingMnemonicPrompt()
	if err != nil {
		return err
	}
	fmt.Println("Generating mnemonic")
	output, err := GetMnemonic(useExisting)
	if err != nil {
		return err
	}
	c.ValidatorMnemonic = output
	c.WithdrawalMnemonic = output

	err = c.GenerateWithdrawalCredentials()
	if err != nil {
		return err
	}
	
	fmt.Println("A mnemonic was generated and stored in node_config.yaml.\n Make sure you don't loose it as you will not be able to recover your keystore if you loose it....")
	return nil
}

func (c *ValidatorCredentials) GenerateWithdrawalCredentials() error {
	promptWithdrawal := promptui.Select{
		Label: "Generate separate withdrawal mnemonic? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, generateVal, err := promptWithdrawal.Run()
	if err != nil {
		return err
	}
	if generateVal == "Yes" {
		c.WithdrawalMnemonic, err = GetMnemonic(false)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ValidatorCredentials) GenerateMnemonicWithoutPrompt() error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}

	fmt.Println("Generating mnemonic")

	useExisting, err := UseExistingMnemonicPrompt()
	if err != nil {
		return err
	}

	output, err := GetMnemonic(useExisting)
	if err != nil {
		return err
	}
	c.ValidatorMnemonic = output
	output, err = GetMnemonic(useExisting)
	if err != nil {
		return err
	}
	c.WithdrawalMnemonic = output

	fmt.Println("A mnemonic was generated and stored in node_config.yaml.\n Make sure you don't loose it as you will not be able to recover your keystore if you loose it....")
	return nil
}

func (c *ValidatorCredentials) GenerateDepositData(details *DepositDetails, numberOfValidators int) error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}

	depositCmd := exec.Command("./bin/network-validator-tool", "deposit-data",
		"--as-json-list",
		"--fork-version", details.ForkVersion,
		"--source-max", fmt.Sprintf("%d", numberOfValidators),
		"--source-min", "0",
		"--amount", details.Amount,
		"--validators-mnemonic", c.ValidatorMnemonic,
		"--withdrawals-mnemonic", c.WithdrawalMnemonic,
	)
	commandOutput, err := depositCmd.Output()
	if err != nil {
		return err
	}
	err = os.WriteFile(details.DepositFileLocation, commandOutput, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (c *ValidatorCredentials) GenerateKeystore(numberOfValidators int, password string) error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}
	nodeConfigs := MustGetNodeConfig()
	keyStoreLocation := nodeConfigs.Keystore.Volume
	walletCmd := exec.Command("./bin/network-validator-tool", "keystores",
		"--insecure",
		"--out-loc", keyStoreLocation,
		"--prysm-pass", password,
		"--source-max", fmt.Sprintf("%d", numberOfValidators),
		"--source-min", "0",
		"--source-mnemonic", c.ValidatorMnemonic,
	)
	err = walletCmd.Run()
	if err != nil {
		return err
	}
	passwdFile := path.Join(keyStoreLocation, "password.txt")
	return os.WriteFile(passwdFile, []byte(password), os.ModePerm)
}

func (c *ValidatorCredentials) GenerateDepositDataWithRange(details *DepositDetails, vRange types.ValidatorRange) error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}

	depositCmd := exec.Command("./bin/network-validator-tool", "deposit-data",
		"--as-json-list",
		"--fork-version", details.ForkVersion,
		"--source-max", fmt.Sprintf("%d", vRange.To),
		"--source-min", fmt.Sprintf("%d", vRange.From),
		"--amount", details.Amount,
		"--validators-mnemonic", c.ValidatorMnemonic,
		"--withdrawals-mnemonic", c.WithdrawalMnemonic,
	)
	commandOutput, err := depositCmd.Output()
	if err != nil {
		return err
	}
	err = os.WriteFile(details.DepositFileLocation, commandOutput, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (c *ValidatorCredentials) GenerateKeystoreWithRange(vRange types.ValidatorRange, password string) error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}
	nodeConfigs := MustGetNodeConfig()
	keyStoreLocation := nodeConfigs.Keystore.Volume
	walletCmd := exec.Command("./bin/network-validator-tool", "keystores",
		"--insecure",
		"--out-loc", keyStoreLocation,
		"--prysm-pass", password,
		"--source-max", fmt.Sprintf("%d", vRange.To),
		"--source-min", fmt.Sprintf("%d", vRange.From),
		"--source-mnemonic", c.ValidatorMnemonic,
	)
	err = walletCmd.Run()
	if err != nil {
		return err
	}
	passwdFile := path.Join(keyStoreLocation, "password.txt")
	return os.WriteFile(passwdFile, []byte(password), os.ModePerm)
}
