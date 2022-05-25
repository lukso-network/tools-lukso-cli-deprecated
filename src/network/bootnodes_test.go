package network

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestBootnodeUpdater_DownloadLatestBootnodesL16Beta(t *testing.T) {
	ba := NewBootnodeUpdater(L16Beta)
	result, err := ba.DownloadLatestBootnodes()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(result)
	assert.Assert(t, len(result) > 0, "bootnodes should be at least 1")
}
