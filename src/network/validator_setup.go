package network

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	types2 "github.com/lukso-network/lukso-cli/src/network/types"
	"github.com/manifoldco/promptui"
	"io"
	"os"
	"os/exec"
	"path"
	"runtime"
)

// CheckAndDownloadValTool checks if validator tool is present or not. If not then download the validator tool according to platform (linux / darwin)
func CheckAndDownloadValTool() error {
	if !FileExists("./bin/network-validator-tool") {
		fmt.Println("downloading network-validator-tool for your system")
		valToolLocation := fmt.Sprintf("https://github.com/lukso-network/network-validator-tools/releases/download/v1.0.0/network-validator-tools-v1.0.0-%s-%s", runtime.GOOS, runtime.GOARCH)
		err := downloadFile(valToolLocation, "./bin/network-validator-tool")
		if err != nil {
			return err
		}
		return os.Chmod("./bin/network-validator-tool", os.ModePerm)
	}
	return nil
}

func GetMnemonic() (string, error) {
	output, err := exec.Command("./bin/network-validator-tool", "mnemonic").Output()
	if err != nil {
		return "", err
	}
	return string(output), err
}

func (valSec *ValidatorCredentials) GenerateMnemonic() error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}

	fmt.Println("Generating mnemonic")

	output, err := GetMnemonic()
	if err != nil {
		return err
	}
	valSec.ValidatorMnemonic = output
	valSec.WithdrawalMnemonic = output

	propmt := promptui.Select{
		Label: "Generate separate withdrawal mnemonic? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, generateVal, err := propmt.Run()
	if err != nil {
		return err
	}
	if generateVal == "Yes" {
		valSec.WithdrawalMnemonic, err = GetMnemonic()
		if err != nil {
			return err
		}
	}

	fmt.Println("A mnemonic was generated and stored in node_config.yaml.\n Make sure you don't loose it as you will not be able to recover your keystore if you loose it....")
	return nil
}

func (valSec *ValidatorCredentials) GenerateDepositData(details *DepositDetails, numberOfValidators int) error {
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
		"--validators-mnemonic", valSec.ValidatorMnemonic,
		"--withdrawals-mnemonic", valSec.WithdrawalMnemonic,
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

func (valSec *ValidatorCredentials) GenerateKeystore(numberOfValidators int, password string) error {
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
		"--source-mnemonic", valSec.ValidatorMnemonic,
	)
	err = walletCmd.Run()
	if err != nil {
		return err
	}
	passwdFile := path.Join(keyStoreLocation, "password.txt")
	return os.WriteFile(passwdFile, []byte(password), os.ModePerm)
}

func (valSec *ValidatorCredentials) GenerateDepositDataWithRange(details *DepositDetails, vRange types2.ValidatorRange) error {
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
		"--validators-mnemonic", valSec.ValidatorMnemonic,
		"--withdrawals-mnemonic", valSec.WithdrawalMnemonic,
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

func (valSec *ValidatorCredentials) GenerateKeystoreWithRange(vRange types2.ValidatorRange, password string) error {
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
		"--source-mnemonic", valSec.ValidatorMnemonic,
	)
	err = walletCmd.Run()
	if err != nil {
		return err
	}
	passwdFile := path.Join(keyStoreLocation, "password.txt")
	return os.WriteFile(passwdFile, []byte(password), os.ModePerm)
}

func (valSec *ValidatorCredentials) DownloadEthereal(ctx context.Context, cli *client.Client) error {
	reader, err := cli.ImagePull(ctx, "docker.io/wealdtech/ethereal:2.7.4", types.ImagePullOptions{})
	if err != nil {
		return err
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}
