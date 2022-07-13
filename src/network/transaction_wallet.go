package network

import "github.com/lukso-network/lukso-cli/src/utils"

func (w *TransactionWallet) CreateNodeRecovery() NodeRecovery {
	if w == nil {
		utils.Coloredln("   warning: transaction wallet is undefined")
		return NodeRecovery{
			TransactionWallet: struct {
				PrivateKey string `json:"privateKey"`
				PublicKey  string `json:"publicKey"`
			}{
				PrivateKey: "",
				PublicKey:  "",
			},
		}
	}
	return NodeRecovery{
		TransactionWallet: struct {
			PrivateKey string `json:"privateKey"`
			PublicKey  string `json:"publicKey"`
		}{
			PrivateKey: w.PrivateKey,
			PublicKey:  w.PublicKey,
		},
	}
}

func (w *TransactionWallet) FromNodeRecovery(nr NodeRecovery) *TransactionWallet {
	w.PublicKey = nr.TransactionWallet.PublicKey
	w.PrivateKey = nr.TransactionWallet.PrivateKey
	return w
}
