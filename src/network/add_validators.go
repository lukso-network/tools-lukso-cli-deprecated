package network

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/manifoldco/promptui"
	"os"
	"strconv"
)

const (
	KeystoreBackupName = "keystore_backup"
	NodeRecoveryBackup = "node_recovery_old_setup.json"
)

type AddValidatorProcess struct {
	configs         *NodeConfigs
	numOfValidators int64
	numOfAdds       int64
	password        string
}

func NewAddValidatorProcess(configs *NodeConfigs, password string) *AddValidatorProcess {
	return &AddValidatorProcess{configs: configs, password: password}
}

func (av *AddValidatorProcess) Add() {
	err := av.setupAddition()
	if err != nil {
		utils.PrintColoredErrorWithReason("couldn't get num of new validators", err)
		return
	}

	err = av.backupKeystore()
	if err != nil {
		utils.PrintColoredErrorWithReason("couldn't recover keystore", err)
		return
	}

	err = av.createNewKeystore()
	if err != nil {
		utils.PrintColoredErrorWithReason("couldn't create new keystore", err)
		err = av.recoverKeystore()
		if err != nil {
			utils.PrintColoredErrorWithReason("couldn't recover keystore", err)
		}
		return
	}

	err = av.cleanupKeystore()
	if err != nil {
		utils.PrintColoredErrorWithReason("couldn't clean up keystore backup files", err)
		return
	}
}

func (av *AddValidatorProcess) setupAddition() error {
	credentials := av.configs.ValidatorCredentials
	numOfValidators := credentials.ValidatorIndexTo - credentials.ValidatorIndexFrom
	fmt.Printf("You currently have %d validators, the index range is [%d, %d] \n", numOfValidators, credentials.ValidatorIndexFrom, credentials.ValidatorIndexTo)
	// set number of validators
	prompt := promptui.Prompt{
		Label: "How many validators do you want to add?",
	}

	numOfAddsString, err := prompt.Run()
	if err != nil {
		return err
	}

	adds, err := strconv.ParseInt(numOfAddsString, 10, 64)
	if err != nil {
		return err
	}

	av.numOfValidators = numOfValidators
	av.numOfAdds = adds
	fmt.Println("You want to add", av.numOfAdds, "new validators. This will result in a total number of", av.numOfAdds+av.numOfValidators, "validators.")
	return nil
}

func (av *AddValidatorProcess) backupKeystore() error {
	fmt.Println("Creating a backup of your current keystore setup...")
	configs := *av.configs
	err := configs.CreateNodeRecovery().SaveWithDestination(NodeRecoveryBackup)
	if err != nil {
		return err
	}

	err = os.Rename(av.configs.Keystore.Volume, KeystoreBackupName)
	if err != nil {
		return err
	}

	fmt.Println("Created recovery file: ", NodeRecoveryBackup, " and moved old keystore to", KeystoreBackupName)
	return nil
}

func (av *AddValidatorProcess) createNewKeystore() error {
	oldCredentials := av.configs.ValidatorCredentials
	oldCredentials.ValidatorIndexTo = av.newNumberOfValidators()
	av.configs.ValidatorCredentials = oldCredentials
	err := av.configs.Save()
	if err != nil {
		return err
	}

	err = av.configs.ValidatorCredentials.GenerateDepositDataWithRange(av.configs.DepositDetails, oldCredentials.ValidatorRange())
	if err != nil {
		return err
	}
	err = av.configs.ValidatorCredentials.GenerateKeystoreWithRange(av.configs.ValidatorCredentials.ValidatorRange(), av.password)
	if err != nil {
		return err
	}
	return nil
}

func (av *AddValidatorProcess) newNumberOfValidators() int64 {
	return av.numOfAdds + av.numOfValidators
}

func (av *AddValidatorProcess) recoverKeystore() error {
	err := os.RemoveAll(av.configs.Keystore.Volume)
	if err != nil {
		return err
	}
	err = os.Rename(KeystoreBackupName, av.configs.Keystore.Volume)
	if err != nil {
		return err
	}
	return nil
}

func (av *AddValidatorProcess) cleanupKeystore() error {
	err := os.RemoveAll(NodeRecoveryBackup)
	if err != nil {
		utils.PrintColoredErrorWithReason("couldn't remove node recovery backup files", err)
		return err
	}
	err = os.RemoveAll(KeystoreBackupName)
	if err != nil {
		utils.PrintColoredErrorWithReason("couldn't remove keystore backup files", err)
		return err
	}
	return nil
}
