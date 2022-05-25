package network

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/src/network/contracts"
	"github.com/lukso-network/lukso-cli/src/network/types"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"github.com/pkg/errors"
	"github.com/wealdtech/go-string2eth"
	"io/ioutil"
	"math/big"
	"strings"
)

func Deposit(depositData string, contractAddress string, privateKey string, rpcEndpoint string) error {
	di, err := loadDepositInfo(depositData)
	if err != nil {
		return err
	}

	err = verifyDepositInfo(di)
	if err != nil {
		return err
	}

	tk, err := wallet.TransactionKeysFromPrivateKey(privateKey)
	if err != nil {
		return err
	}

	fmt.Printf("dialing into rpc endpoint %s........ ", rpcEndpoint)
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return err
	}
	fmt.Println("success")

	balance, err := client.BalanceAt(context.Background(), tk.PublicKey, nil)
	if err != nil {
		return err
	}
	fmt.Printf("Balance of transaction_key(%s): %s\n", tk.PublicKey, string2eth.WeiToString(balance, true))
	contract, err := contracts.NewEth2Deposit(common.HexToAddress(contractAddress), client)
	if err != nil {
		return err
	}

	for k, d := range di {
		opts := createTransactionOpts(client, &tk)
		opts.Value = new(big.Int).Mul(new(big.Int).SetUint64(d.Amount), big.NewInt(1000000000))
		opts.Context = context.Background()

		fmt.Printf("Creating %s deposit for key: %d\n", string2eth.WeiToString(opts.Value, true), d.PublicKey)

		var depositDataRoot [32]byte
		copy(depositDataRoot[:], d.DepositDataRoot)
		fmt.Printf("Waiting for transaction no %d to be mined ....", k)
		signedTx, err := contract.Deposit(opts, d.PublicKey, d.WithdrawalCredentials, d.Signature, depositDataRoot)
		if err != nil {
			return err
		}
		receipt, err := bind.WaitMined(context.Background(), client, signedTx)
		if err != nil {
			return err
		}
		fmt.Println("mined at: ", receipt.BlockNumber)
	}

	return nil
}

func loadDepositInfo(input string) ([]*types.DepositInfo, error) {
	var err error
	var data []byte
	// Input could be JSON or a path to JSON
	switch {
	case strings.HasPrefix(input, "{"):
		// Looks like JSON
		data = []byte("[" + input + "]")
	case strings.HasPrefix(input, "["):
		// Looks like JSON array
		data = []byte(input)
	default:
		// Assume it's a path to JSON
		data, err = ioutil.ReadFile(input)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find deposit data file")
		}
		if data[0] == '{' {
			data = []byte("[" + string(data) + "]")
		}
	}

	depositInfo, err := types.DepositInfoFromJSON(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to obtain deposit information")
	}
	if len(depositInfo) == 0 {
		return nil, errors.New("no deposit information supplied")
	}
	return depositInfo, nil
}

// TODO needs proper verification
func verifyDepositInfo(depositInfo []*types.DepositInfo) error {
	for k, d := range depositInfo {
		if len(d.PublicKey) == 0 {
			return fmt.Errorf("no public key for deposit %d", k)
		}
		if len(d.DepositDataRoot) == 0 {
			return fmt.Errorf("no data root for deposit %d", k)
		}
		if len(d.Signature) == 0 {
			return fmt.Errorf("no signature for deposit %d", k)
		}
		if len(d.WithdrawalCredentials) == 0 {
			return fmt.Errorf("no ithdrawal credentials for deposit %d", k)
		}
		//		if len(contract.forkVersion) != 0 && len(depositInfo[i].ForkVersion) != 0 {
		//			cli.Assert(bytes.Equal(depositInfo[i].ForkVersion, contract.forkVersion), quiet, fmt.Sprintf("Incorrect fork version for deposit %d (expected %#x, found %#x)", i, contract.forkVersion, depositInfo[i].ForkVersion))
		//		}
		//		cli.Assert(depositInfo[i].Amount >= 1000000000, quiet, fmt.Sprintf("Deposit too small for deposit %d", i))
		//		cli.Assert(depositInfo[i].Amount <= 32000000000 || beaconDepositAllowExcessiveDeposit, quiet, fmt.Sprintf(`Deposit more than 32 Ether for deposit %d.  Any amount above 32 Ether that is deposited will not count towards the validator's effective balance, and is effectively wasted.
		//
		//If you really want to do this use the --allow-excessive-deposit option.`, i))
		//
		//		cli.Assert(beaconDepositAllowOldData || depositInfo[i].Version >= contract.minVersion, quiet, `Data generated by ethdo is old and possibly inaccurate.  This means you need to upgrade your version of ethdo (or you are sending your deposit to the wrong contract or network); please do so by visiting https://github.com/wealdtech/ethdo and following the installation instructions there.  Once you have done this please regenerate your deposit data and try again.
		//
		//If you are *completely sure* you know what you are doing, you can use the --allow-old-data option to carry out this transaction.  Otherwise, please seek support to ensure you do not lose your Ether.`)
		//		cli.Assert(beaconDepositAllowNewData || depositInfo[i].Version <= contract.maxVersion, quiet, `Data generated by ethdo is newer than supported.  This means you need to upgrade your version of ethereal (or you are sending your deposit to the wrong contract or network); please do so by visiting https://github.com/wealdtech/ethereal and following the installation instructions there.  Once you have done this please try again.
		//
		//If you are *completely sure* you know what you are doing, you can use the --allow-new-data option to carry out this transaction.  Otherwise, please seek support to ensure you do not lose your Ether.`)
	}
	return nil
}

func createTransactionOpts(client *ethclient.Client, tk *wallet.TransactionKeys) *bind.TransactOpts {
	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), tk.PublicKey)
	if err != nil {
		panic(err)
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(tk.PrivateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(160000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}