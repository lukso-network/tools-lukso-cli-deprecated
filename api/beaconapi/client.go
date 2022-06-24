package beaconapi

import "github.com/lukso-network/lukso-cli/api/http"

// DOCS for Beacon API https://ethereum.github.io/beacon-APIs

type BeaconClient struct {
	client http.HttpClient
}

func NewBeaconClient(baseUrl string) BeaconClient {
	return BeaconClient{
		http.NewHttpClient(baseUrl),
	}
}
