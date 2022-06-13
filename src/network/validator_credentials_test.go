package network

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatorCredentialsFromNodeRecovery(t *testing.T) {
	defer CleanUpRecoveryTest()
	nr := NodeRecovery{
		ValidatorMnemonic:  "vm",
		WithdrawalMnemonic: "wm",
		KeystoreIndexFrom:  10,
		KeystoreIndexTo:    18,
	}

	actualNr := new(ValidatorCredentials).FromNodeRecovery(nr).CreateNodeRecovery()

	require.Equal(t, nr.KeystoreIndexFrom, actualNr.KeystoreIndexFrom)
	require.Equal(t, nr.KeystoreIndexTo, actualNr.KeystoreIndexTo)
	require.Equal(t, nr.WithdrawalMnemonic, actualNr.WithdrawalMnemonic)
	require.Equal(t, nr.ValidatorMnemonic, actualNr.ValidatorMnemonic)
}
