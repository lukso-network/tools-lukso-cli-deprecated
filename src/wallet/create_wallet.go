package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	passwordLength = 32
)

type WalletInfo struct {
	PubKey  string
	PrivKey string
}

func CreateWallet(targetDirectory string, password string, label string) (*WalletInfo, error) {
	if password == "" {
		password = CreateRandomPassword()
	}
	store := keystore.NewKeyStore(targetDirectory, keystore.StandardScryptN, keystore.StandardScryptP)

	a, err := store.NewAccount(password)
	if err != nil {
		return nil, err
	}

	filename := a.URL.String()
	passwordFilename := strings.Replace(a.URL.Path, a.URL.Scheme, "", 1)

	if label != "" {

		if targetDirectory == "" {
			filename = fmt.Sprintf("%v.json", label)
		} else {
			filename = fmt.Sprintf("%v/%v.json", targetDirectory, label)
		}
		err = os.Rename(strings.Replace(a.URL.Path, a.URL.Scheme, "", 1), filename)
		if err != nil {
			return nil, err
		}

		passwordFilename = label
	}

	// write password file
	if targetDirectory == "" {
		err := ioutil.WriteFile(fmt.Sprintf("%v_password.txt", passwordFilename), []byte(password), os.ModePerm)
		if err != nil {
			return nil, err
		}
	} else {
		err := ioutil.WriteFile(fmt.Sprintf("%v/%v_password.txt", targetDirectory, passwordFilename), []byte(password), os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	keyJSON, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("couldn't read wallet file %v %v", filename, err.Error())
	}

	key, err := keystore.DecryptKey(keyJSON, password)

	if err != nil {
		return nil, err
	}
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)

	return &WalletInfo{
		PubKey:  strings.ToLower(key.Address.String()),
		PrivKey: strings.Replace(hexutil.Encode(privateKeyBytes), "0x", "", 1),
	}, nil
}

func CreateRandomPassword() string {
	return randStringBytes(passwordLength)
}

func randStringBytes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
