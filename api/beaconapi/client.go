package beaconapi

import "github.com/lukso-network/lukso-cli/api/http"

const DefaultBeaconAPIEndpoint = "34.90.85.198:3500"

type BeaconClient struct {
	client http.HttpClient
}

func NewBeaconClient(baseUrl string) BeaconClient {
	return BeaconClient{
		http.NewHttpClient(baseUrl),
	}
}
