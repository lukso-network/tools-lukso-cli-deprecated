package network

import (
	"errors"
	"fmt"
	"os/exec"
)

func runDockerServices(serviceList ...string) error {
	dockerCommand := []string{"docker-compose", "up", "-d"}
	dockerCommand = append(dockerCommand, serviceList...)
	fmt.Println("You may need to provide super user (sudo) password to run docker (if needed)")
	command := exec.Command("sudo", dockerCommand...)
	if err := command.Run(); err != nil {
		return fmt.Errorf("found error while running docker. Make sure your docker is running. %s", err)
	}
	return nil
}

func DownDockerServices() error {
	command := exec.Command("sudo", "docker-compose", "down")
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func StartArchNode() error {
	config, err := GetLoadedNodeConfigs()
	if err != nil {
		return err
	}
	configDirName, err := config.getConfigPath()
	if err != nil {
		return err
	}
	if configDirName == "" || !FileExists(configDirName) {
		return errors.New("config files are not present. Can't start docker containers")
	}
	return runDockerServices("init-geth", "geth", "prysm_beacon", "eth2stats-client")
}

func StartValidatorNode() error {
	config, err := GetLoadedNodeConfigs()
	if err != nil {
		return err
	}
	configDirName, err := config.getConfigPath()
	if err != nil {
		return err
	}
	keystorePath, err := config.getKeyStorePath()
	if err != nil {
		return err
	}
	if configDirName == "" || !FileExists(configDirName) {
		return errors.New("config files are not present. Can't start docker containers")
	}
	if !FileExists(keystorePath) {
		return errors.New("keystore path is invalid")
	}
	return runDockerServices("prysm_validator")
}
