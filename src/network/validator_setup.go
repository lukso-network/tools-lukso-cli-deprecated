package network

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/manifoldco/promptui"
	"io"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
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

func (valSec *ValidatorSecrets) GenerateMnemonic() error {
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
	return nil
}

func (valSec *ValidatorSecrets) GenerateDepositData(numberOfValidators int) error {

	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}

	depositCmd := exec.Command("./bin/network-validator-tool", "deposit-data",
		"--as-json-list",
		"--fork-version", valSec.ForkVersion,
		"--source-max", fmt.Sprintf("%d", numberOfValidators),
		"--source-min", "0",
		"--validators-mnemonic", valSec.ValidatorMnemonic,
		"--withdrawals-mnemonic", valSec.WithdrawalMnemonic,
	)
	commandOutput, err := depositCmd.Output()
	if err != nil {
		return err
	}
	err = os.WriteFile(valSec.Deposit.DepositFileLocation, commandOutput, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (valSec *ValidatorSecrets) GenerateWallet(numberOfValidators int, password string) error {
	err := CheckAndDownloadValTool()
	if err != nil {
		return err
	}
	nodeConfigs, err := GetLoadedNodeConfigs()
	if err != nil {
		return err
	}
	keyStoreLocation, err := nodeConfigs.GetKeyStorePath()
	if err != nil {
		return err
	}
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

func PullEtherealImage(ctx context.Context, client *client.Client) error {
	reader, err := client.ImagePull(ctx, "wealdtech/ethereal", types.ImagePullOptions{})
	if err != nil {
		return err
	}

	defer reader.Close()
	_, err = io.Copy(os.Stdout, reader)
	return err
}

func (valSec *ValidatorSecrets) DownloadEthereal(ctx context.Context, cli *client.Client) error {
	reader, err := cli.ImagePull(ctx, "docker.io/wealdtech/ethereal", types.ImagePullOptions{})
	if err != nil {
		return err
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)
	return nil
}

func (valSec *ValidatorSecrets) DoDeposit(ctx context.Context, data *DepositData, cli *client.Client) error {
	depData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "wealdtech/ethereal",
		Cmd: []string{"beacon", "deposit",
			"--allow-unknown-contract", strconv.FormatBool(valSec.Deposit.Force),
			"--address", valSec.Deposit.ContractAddress,
			"--connection", valSec.Eth1Data.RPCEndPoint,
			"--value", valSec.Deposit.Amount,
			"--from", valSec.Eth1Data.WalletAddress,
			"--privatekey", valSec.Eth1Data.WalletPrivKey,
			"--data", string(depData),
		},
		Tty: false,
	}, nil, nil, nil, "")
	if err != nil {
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("Sending deposit for validator", data.PubKey)
	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return err
}

func (valSec *ValidatorSecrets) SendDepositTxn() error {
	if valSec.Eth1Data.WalletAddress == "" {
		prompt := promptui.Prompt{
			Label: "Enter valid wallet address",
		}
		eth1Address, err := prompt.Run()
		if err != nil {
			return err
		}
		valSec.Eth1Data.WalletAddress = eth1Address
	}
	if valSec.Eth1Data.WalletPrivKey == "" {
		prompt := promptui.Prompt{
			Label: "Enter valid wallet private key",
			Mask:  '*',
		}
		eth1PrivateKey, err := prompt.Run()
		if err != nil {
			return err
		}
		valSec.Eth1Data.WalletPrivKey = eth1PrivateKey
	}
	depositData, err := ParseDepositDataFromFile(valSec.Deposit.DepositFileLocation)
	if err != nil {
		return err
	}
	dockerClient, err := GetDockerClient()
	if err != nil {
		return err
	}
	err = valSec.DownloadEthereal(context.Background(), dockerClient)
	if err != nil {
		return err
	}
	for _, data := range depositData {
		err = valSec.DoDeposit(context.Background(), data, dockerClient)
		if err != nil {
			return err
		}
	}
	return nil
}
