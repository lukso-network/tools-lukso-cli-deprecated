package network

func (w *TransactionWallet) FromNodeRecovery(nr NodeRecovery) *TransactionWallet {
	w.PublicKey = nr.TransactionWallet.PublicKey
	w.PrivateKey = nr.TransactionWallet.PrivateKey
	return w
}

func (w *TransactionWallet) IsEmpty() bool {
	return w.PublicKey == "" || w.PrivateKey == ""
}
