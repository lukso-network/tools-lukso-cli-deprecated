package utils

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/src/wallet"
	"math/big"
	"time"
)

func SendTransaction(tx *types.Transaction, client *ethclient.Client, limit time.Duration, timeout time.Duration) bool {
	mined := WaitForTransaction(client, tx.Hash(), limit, timeout)
	if mined {
		fmt.Sprintf("%s mined\n", tx.Hash().Hex())
	} else {
		fmt.Sprintf("%s submitted but not mined\n", tx.Hash().Hex())
	}

	return mined
}

// WaitForTransaction waits for the transaction to be mined, or for the limit to expire
func WaitForTransaction(client *ethclient.Client, txHash common.Hash, limit time.Duration, timeout time.Duration) bool {
	start := time.Now()
	first := true
	for limit == 0 || time.Since(start) < limit {
		if !first {
			time.Sleep(5 * time.Second)
		} else {
			first = false
		}
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_, pending, err := client.TransactionByHash(ctx, txHash)
		if err == nil && !pending {
			return true
		}
	}
	return false
}

func CreateTransactionOpts(client *ethclient.Client, tk *wallet.TransactionKeys) *bind.TransactOpts {
	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), tk.PublicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(tk.PrivateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
