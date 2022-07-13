package network

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatorCredentialsFromNodeRecovery(t *testing.T) {
	defer CleanUpRecoveryTest()

	nr := NodeRecovery{
		ValidatorCredentials: struct {
			ValidatorMnemonic  string `json:"validatorMnemonic"`
			WithdrawalMnemonic string `json:"withdrawalMnemonic"`
			KeystoreIndexFrom  int64  `json:"keystoreIndexFrom"`
			KeystoreIndexTo    int64  `json:"keystoreIndexTo"`
		}{
			ValidatorMnemonic:  "vm",
			WithdrawalMnemonic: "wm",
			KeystoreIndexFrom:  10,
			KeystoreIndexTo:    18,
		},
	}
	actualNr := new(ValidatorCredentials).FromNodeRecovery(nr).CreateNodeRecovery()

	require.Equal(t, nr.ValidatorCredentials.KeystoreIndexFrom, actualNr.ValidatorCredentials.KeystoreIndexFrom)
	require.Equal(t, nr.ValidatorCredentials.KeystoreIndexTo, actualNr.ValidatorCredentials.KeystoreIndexTo)
	require.Equal(t, nr.ValidatorCredentials.WithdrawalMnemonic, actualNr.ValidatorCredentials.WithdrawalMnemonic)
	require.Equal(t, nr.ValidatorCredentials.ValidatorMnemonic, actualNr.ValidatorCredentials.ValidatorMnemonic)
}
