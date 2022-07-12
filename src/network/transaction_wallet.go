package network

func (w *TransactionWallet) CreateNodeRecovery() NodeRecovery {
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
