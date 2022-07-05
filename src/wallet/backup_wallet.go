package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lukso-network/lukso-cli/src/network"
	"io/ioutil"
	"strings"
)

func BackupWallet(target, filename, passwordFilename string) (*WalletInfo, error) {
	keyJSON, err := ioutil.ReadFile(fmt.Sprintf("%v/%v.json", target, filename))
	if err != nil {
		return nil, fmt.Errorf("couldn't read wallet file %v %v", filename, err.Error())
	}

	password, err := ioutil.ReadFile(fmt.Sprintf("%v/%v_password.txt", target, passwordFilename))
	if err != nil {
		return nil, fmt.Errorf("couldn't read wallet password file %v %v", passwordFilename, err.Error())
	}

	key, err := keystore.DecryptKey(keyJSON, string(password))
	if err != nil {
		return nil, fmt.Errorf("could not decrypt keystore: %v", err.Error())
	}
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)

	return &WalletInfo{
		PubKey:  strings.ToLower(key.Address.String()),
		PrivKey: strings.Replace(hexutil.Encode(privateKeyBytes), "0x", "", 1),
	}, nil
}

func (tw network.TransactionWallet) CreateWalletRecovery() network.NodeRecovery {
	return network.NodeRecovery{}
}
