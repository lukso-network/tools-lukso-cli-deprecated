package network

import (
	"bytes"
	"fmt"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"os"
	"os/exec"
	"runtime"
)

type WalletChunk struct {
	Start    int
	End      int
	Name     string
	IsSingle bool
}

var BetaConfigName = "beta"

type GenesisCreatorConfig struct {
	TargetDirectory string
	Mnemonic        string
	WalletChunks    []WalletChunk
}

var betaTestNetworkGenesisConfig = GenesisCreatorConfig{
	TargetDirectory: "beta_genesis_wallets",
	WalletChunks:    []WalletChunk{},
}

func CreateBetaConfig(mnemonic string) GenesisCreatorConfig {
	betaTestNetworkGenesisConfig.Mnemonic = mnemonic

	every := 10
	// first 30 nodes
	for i := 0; i < 30; i++ {
		betaTestNetworkGenesisConfig.WalletChunks = append(betaTestNetworkGenesisConfig.WalletChunks, WalletChunk{
			Start:    i * every,
			End:      i*every + every,
			Name:     fmt.Sprintf("validator%d", i+1),
			IsSingle: false,
		})
	}
	betaTestNetworkGenesisConfig.WalletChunks = append(betaTestNetworkGenesisConfig.WalletChunks, WalletChunk{
		Start:    301,
		End:      400,
		Name:     "beta",
		IsSingle: true,
	})

	return betaTestNetworkGenesisConfig
}

func CreateValidatorWallets(mnemonic string, configType string) error {
	var config GenesisCreatorConfig
	switch configType {
	case BetaConfigName:
		config = CreateBetaConfig(mnemonic)
	default:
		return fmt.Errorf("config type not found")
	}

	fmt.Println("Mnemonic", config.Mnemonic)
	fmt.Println("TargetDirectory", config.TargetDirectory)
	for _, c := range config.WalletChunks {
		if c.IsSingle {
			if err := createSingleWallet(config.TargetDirectory, config.Mnemonic, c); err != nil {
				fmt.Println(err.Error())
				return err
			}
		} else {
			if err := createWallet(config.TargetDirectory, config.Mnemonic, c); err != nil {
				fmt.Println(err.Error())
				return err
			}
		}
	}

	return nil
}

var commandTemplateDarwin = "./bin/eth2-val-tools-Darwin-x86_64"
var commandTemplateLinux = "./bin/eth2-val-tools-Linux-x86_64"

func createWallet(dir string, mnemonic string, chunk WalletChunk) error {
	directory := fmt.Sprintf("%s/%s", dir, chunk.Name)
	return executeCommand(directory, chunk.End, chunk.Start, mnemonic)
}

func createSingleWallet(dir string, mnemonic string, chunk WalletChunk) error {
	for i := chunk.Start; i <= chunk.End; i++ {
		name := fmt.Sprintf("beta_%d", i-chunk.Start+1)
		fmt.Println("Creating", name)
		directory := fmt.Sprintf("%s/%s/%s", dir, chunk.Name, name)
		err := executeCommand(directory, i+1, i, mnemonic)
		if err != nil {
			return err
		}

	}
	return nil
}

func executeCommand(directory string, end int, start int, mnemonic string) error {
	password := wallet.CreateRandomPassword()
	command := commandTemplateLinux
	switch runtime.GOOS {
	case "darwin":
		command = commandTemplateDarwin
	}
	cmd :=
		exec.Command(command,
			"keystores",
			"--insecure",
			"--out-loc", directory,
			"--prysm-pass", password,
			"--source-max", fmt.Sprintf("%d", end),
			"--source-min", fmt.Sprintf("%d", start),
			"--source-mnemonic", fmt.Sprintf("%s", mnemonic),
		)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	fmt.Println("Result: " + out.String())

	file, err := os.Create(fmt.Sprintf("%s/password.txt", directory))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(password)

	if err != nil {
		return err
	}

	return nil
}
