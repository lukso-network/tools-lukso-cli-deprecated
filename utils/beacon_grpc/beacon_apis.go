package beacon_grpc

type BeaconAPIs interface {
	GetValidatorStatus(pubKey []byte) (string, error)
}
