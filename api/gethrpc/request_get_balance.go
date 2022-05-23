package gethrpc

func (client *Instance) GetBalance(address string) (int64, error) {
	return client.RequestInt64("eth_getBalance", address)
}
