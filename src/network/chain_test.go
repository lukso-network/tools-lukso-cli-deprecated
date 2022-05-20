package network

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsChainSupported(t *testing.T) {
	if IsChainSupported(L16) {
		t.Error("L16 is not supported yet")
		return
	}

	if IsChainSupported(MainNet) {
		t.Error("MainNet is not supported yet")
		return
	}

	if !IsChainSupported(L16Beta) {
		t.Error("L16Beta  should be supported ")
		return
	}
}

func TestGetChainByString(t *testing.T) {
	require.Equal(t, L16, GetChainByString("l16"))
	require.Equal(t, L16, GetChainByString("L16"))
	require.Equal(t, L16Beta, GetChainByString("L16Beta"))
	require.Equal(t, L16Beta, GetChainByString("l16beta"))
	require.Equal(t, L16Beta, GetChainByString("l16Beta"))
	require.Equal(t, L16Beta, GetChainByString("L16beta"))
	require.Equal(t, MainNet, GetChainByString("Mainnet"))
	require.Equal(t, MainNet, GetChainByString("mainnet"))
}
