package network

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPublicIP(t *testing.T) {
	t.Parallel()
	t.Run("test public ip address", func(t *testing.T) {
		_, err := getPublicIP()
		require.NoError(t, err)
	})
}
