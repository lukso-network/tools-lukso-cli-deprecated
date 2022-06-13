package network

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSaveAndLoadNodeRecovery(t *testing.T) {
	defer CleanUpRecoveryTest()
	nr := NodeRecovery{
		ValidatorMnemonic:  "vm",
		WithdrawalMnemonic: "wm",
		KeystoreIndexFrom:  1,
		KeystoreIndexTo:    100,
	}

	err := nr.Save()
	require.NoError(t, err)

	actualNr, err := LoadNodeRecovery(NodeRecoveryFileLocation)
	require.NoError(t, err)

	require.Equal(t, nr.KeystoreIndexFrom, actualNr.KeystoreIndexFrom)
	require.Equal(t, nr.KeystoreIndexTo, actualNr.KeystoreIndexTo)
	require.Equal(t, nr.WithdrawalMnemonic, actualNr.WithdrawalMnemonic)
	require.Equal(t, nr.ValidatorMnemonic, actualNr.ValidatorMnemonic)
}

func CleanUpRecoveryTest() {
	_ = os.Remove(NodeRecoveryFileLocation)
}
