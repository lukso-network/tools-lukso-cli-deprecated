package network

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/src/network/contracts"
	"github.com/lukso-network/lukso-cli/src/network/types"
	"github.com/lukso-network/lukso-cli/src/utils"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"github.com/pkg/errors"
	"io/ioutil"
	"math/big"
	"strings"
)

func Deposit(
	contractDepositEvents *DepositEvents,
	depositData string,
	contractAddress string,
	privateKey string,
	rpcEndpoint string,
	maxGasFee int64,
	priorityFee int64,
	dry bool) (int, error) {
	di, err := loadDepositInfo(depositData)
	if err != nil {
		return -1, err
	}

	err = verifyDepositInfo(di)
	if err != nil {
		return -1, err
	}

	tk, err := wallet.TransactionKeysFromPrivateKey(privateKey)
	if err != nil {
		return -1, err
	}
	fmt.Println(tk)

	fmt.Printf("dialing into rpc endpoint %s........ ", rpcEndpoint)
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return -1, err
	}
	fmt.Println("success")

	balance, err := client.BalanceAt(context.Background(), tk.PublicKey, nil)
	if err != nil {
		return -1, err
	}

	fmt.Printf("Balance of transaction_key(%s): %s\n", tk.PublicKey, utils.WeiToString(balance, true))
	contract, err := contracts.NewEth2Deposit(common.HexToAddress(contractAddress), client)
	if err != nil {
		return -1, err
	}

	totalDeposited := 0
	for k, d := range di {
		pubKey := utils.MaybeAddHexPrefix(common.Bytes2Hex(d.PublicKey))
		amount := new(big.Int).Mul(new(big.Int).SetUint64(d.Amount), big.NewInt(1000000000))

		fmt.Printf("Creating %s deposit for key: %s (PriorityFee/MaxFee %d/%d)\n", utils.WeiToString(amount, true), common.Bytes2Hex(d.PublicKey), priorityFee, maxGasFee)

		totalDepositedAmount := contractDepositEvents.Amount(pubKey)

		if totalDepositedAmount > 0 {
			fmt.Println("Validator has already a deposit with amount: \n", totalDepositedAmount)
			//continue
		}

		opts, err := createTransactionOpts(client, &tk, priorityFee, maxGasFee)
		if err != nil {
			fmt.Println("The transaction failed, reason: ", err.Error())
			continue
		}
		opts.Value = amount
		opts.Context = context.Background()

		var depositDataRoot [32]byte
		copy(depositDataRoot[:], d.DepositDataRoot)
		fmt.Printf("Waiting for transaction no %d to be mined ....", k)
		if dry {
			fmt.Println(" transaction not transmitted - this is a dry run")
			continue
		}
		signedTx, err := contract.Deposit(opts, d.PublicKey, d.WithdrawalCredentials, d.Signature, depositDataRoot)
		if err != nil {
			fmt.Println("The transaction failed, reason: ", err.Error())
			continue
		}

		receipt, err := bind.WaitMined(context.Background(), client, signedTx)
		if err != nil {
			fmt.Println("The transaction failed, reason: ", err.Error())
			continue
		}
		fmt.Println("mined at: ", receipt.BlockNumber)
		totalDeposited++
	}

	return totalDeposited, nil
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
	}
	return nil
}

func createTransactionOpts(client *ethclient.Client, tk *wallet.TransactionKeys, maxFee int64, priorityFee int64) (*bind.TransactOpts, error) {
	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), tk.PublicKey)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(tk.PrivateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(160000) // in units
	auth.GasTipCap = big.NewInt(priorityFee)
	auth.GasFeeCap = big.NewInt(maxFee)
	//auth.GasPrice = big.NewInt(gasPrice)

	return auth, err
}
