package network

import (
	"fmt"
	"github.com/lukso-network/lukso-cli/api/beaconapi"
	"github.com/lukso-network/lukso-cli/src/utils"
	"net/http"
)

func DescribeValidatorKey(keys []string, contractAddress string, executionApi string, consensusApi string, depositEvents *DepositEvents) (err error) {
	fmt.Println("........................................................................................................................................................................")
	utils.ColoredPrintln("Number of validators:", len(keys))
	fmt.Println("Configuration")
	fmt.Println("........................................................................................................................................................................")
	utils.ColoredPrintln("Consensus Api:", consensusApi)
	utils.ColoredPrintln("Execution Api:", executionApi)
	utils.ColoredPrintln("Contract Address:", contractAddress)
	fmt.Println("........................................................................................................................................................................")

	beaconClient := beaconapi.NewBeaconClient(consensusApi)
	fmt.Println("Getting all deposits from contract....")

	if depositEvents == nil {
		e, err := NewDepositEvents(contractAddress, executionApi)
		if err != nil {
			return err
		}
		depositEvents = &e
	}

	for _, k := range keys {
		key := utils.MaybeAddHexPrefix(k)
		fmt.Printf("Checking state of validator key: %v.......", key)
		amount := depositEvents.Amount(key)
		if amount == 0 {
			fmt.Println("   not deposited yet")
			fmt.Println("")
			continue
		}

		state, status, err := beaconClient.ValidatorState(key, -1)
		if status == http.StatusNotFound {
			fmt.Println("  is pending")
			fmt.Println("")
			continue
		}
		if err != nil {
			return err
		}
		fmt.Println("s")
		utils.ColoredPrintln("ValidatorKey", state.Data.Validator.Pubkey)
		utils.ColoredPrintln("Index:", state.Data.Index)
		utils.ColoredPrintln("Status:", state.Data.Status)
		utils.ColoredPrintln("Balance:", state.Data.Balance)
		utils.ColoredPrintln("Effective Balance:", state.Data.Validator.EffectiveBalance)
		utils.ColoredPrintln("Activation Epoch:", state.Data.Validator.ActivationEpoch)
		utils.ColoredPrintln("Activation Eligibility Epoch:", state.Data.Validator.ActivationEligibilityEpoch)
		utils.ColoredPrintln("Exit Epoch:", state.Data.Validator.ExitEpoch)
		utils.ColoredPrintln("Withdrawable Epoch:", state.Data.Validator.WithdrawableEpoch)
		utils.ColoredPrintln("Withdrawal Credentials", state.Data.Validator.WithdrawalCredentials)
		utils.ColoredPrintln("Is Slashed:", state.Data.Validator.Slashed)
	}

	return nil
}
