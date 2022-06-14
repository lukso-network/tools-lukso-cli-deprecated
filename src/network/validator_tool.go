package network

import (
	"fmt"
	"os"
	"os/exec"
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
