package wallet

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"strings"
)

type TransactionKeys struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  common.Address
}

func TransactionKeysFromPrivateKey(privateKeyHex string) (TransactionKeys, error) {
	tk := TransactionKeys{}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return tk, err
	}
	tk.PrivateKey = privateKey

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return tk, fmt.Errorf("unable to create public key from private key")
	}

	tk.PublicKey = crypto.PubkeyToAddress(*publicKeyECDSA)

	return tk, nil
}

func KeyFromWalletAndPasswordFile(keystoreUTCPath string, password string) (*keystore.Key, error) {
	keyJSON, err := ioutil.ReadFile(keystoreUTCPath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read wallet file %v %v", keystoreUTCPath, err.Error())
	}

	key, err := keystore.DecryptKey(keyJSON, password)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func ReadPasswordFile(passwordFile string) (string, error) {
	password, err := ioutil.ReadFile(passwordFile)
	if err != nil {
		return "", fmt.Errorf("couldn't read wallet file %v %v", passwordFile, err.Error())
	}

	return string(password), nil
}

func PrivateKeyFromKey(key *keystore.Key) string {
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)
	return hexutil.Encode(privateKeyBytes)
}

func PublicKeyFromKey(k *keystore.Key) string {
	return k.Address.String()
}

func PrivateKeyFromHex(privateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
}
