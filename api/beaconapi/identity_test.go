package beaconapi

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

const L16Endpoint = "https://beacon.l16.lukso.network"

func TestBeaconClient_Identity(t *testing.T) {
	client := NewBeaconClient(L16Endpoint)

	response, err := client.Identity()

	assert.NilError(t, err, "must return no error when calling identity")

	fmt.Println(response)
}
