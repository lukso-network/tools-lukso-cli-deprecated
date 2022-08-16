package network

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
	"os"
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

func GetMnemonic(existing bool) (string, error) {
	if existing {
		existingMnemonicVal, err := getExistingMnemonic()
		if err != nil {
			return "", err
		}
		return existingMnemonicVal, err
	}
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		errors.Wrap(err, "cannot get 256 bits of entropy")
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		errors.Wrap(err, "cannot get 256 bits of entropy")
	}
	return mnemonic, err
}

func getExistingMnemonic() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter your existing mnemonic",
	}
	existingMnemonicVal, err := prompt.Run()
	if err != nil {
		cobra.CompErrorln(err.Error())
		return "", err
	}
	return existingMnemonicVal, nil
}

func UseExistingMnemonicPrompt() (bool, error) {
	promptExisting := promptui.Select{
		Label: "Use existing mnemonic?? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, existingVal, err := promptExisting.Run()
	if err != nil {
		return false, err
	}
	if existingVal == "Yes" {
		return true, nil
	}
	return false, nil
}
