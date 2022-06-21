package network

import (
	"errors"
	"fmt"
	"os/exec"
)

func runDockerServices(serviceList ...string) error {
	checkDockerIsRunning()
	dockerCommand := []string{"docker-compose", "up", "-d"}
	dockerCommand = append(dockerCommand, serviceList...)
	fmt.Println("You may need to provide super user (sudo) password to run docker (if needed)")
	command := exec.Command("sudo", dockerCommand...)
	if commandOutput, err := command.CombinedOutput(); err != nil {
		return fmt.Errorf("error code: %s. %s", err, string(commandOutput))
	}
	return nil
}

func DownDockerServices() error {
	checkDockerIsRunning()
	command := exec.Command("sudo", "docker-compose", "down")
	if cmdOutput, err := command.CombinedOutput(); err != nil {
		fmt.Println(string(cmdOutput))
		return err
	}
	return nil
}

func StartArchNode() error {
	config := MustGetNodeConfig()
	configDirName := config.Configs.Volume
	if configDirName == "" || !FileExists(configDirName) {
		return errors.New("config files are not present. Can't start docker containers")
	}
	return runDockerServices("init-geth", "geth", "prysm_beacon", "eth2stats-client")
}

func StartValidatorNode() error {
	config := MustGetNodeConfig()
	configDirName := config.Configs.Volume
	keystorePath := config.Keystore.Volume
	if configDirName == "" || !FileExists(configDirName) {
		return errors.New("config files are not present. Can't start docker containers")
	}
	if !FileExists(keystorePath) {
		return errors.New("keystore path is invalid")
	}
	return runDockerServices("prysm_validator")
}

func StopValidatorNode() error {
	command := exec.Command("sudo", "docker-compose", "stop", "prysm_validator")
	if cmdOutput, err := command.CombinedOutput(); err != nil {
		fmt.Println(string(cmdOutput))
		return err
	}
	return nil
}

func checkDockerIsRunning() error {
	cmd := exec.Command("docker", "version", ">", "/dev/null", "2>&1")
	err := cmd.Run()
	if err != nil {
		return errors.New("the Docker daemon is not running. Has Docker been installed")
	}
	return nil
}
