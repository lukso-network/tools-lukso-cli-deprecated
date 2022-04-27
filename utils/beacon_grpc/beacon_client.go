package beacon_grpc

import (
	"context"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BeaconClient struct {
	conn            *grpc.ClientConn
	validatorClient ethpb.BeaconNodeValidatorClient
}

func NewBeaconClient(endpoint string) (*BeaconClient, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// Maximum receive value 128 MB
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(128 * 1024 * 1024)),
	}
	conn, err := grpc.Dial(endpoint, dialOpts...)
	if err != nil {
		return nil, err
	}

	return &BeaconClient{
		conn:            conn,
		validatorClient: ethpb.NewBeaconNodeValidatorClient(conn),
	}, nil
}

func (bc *BeaconClient) GetValidatorStatus(pubKey []byte) (string, error) {
	status, err := bc.validatorClient.ValidatorStatus(context.Background(), &ethpb.ValidatorStatusRequest{PublicKey: pubKey})
	if err != nil {
		return "", err
	}
	return status.GetStatus().String(), nil
}
