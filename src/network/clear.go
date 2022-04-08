package network

import "os"

func Clear() error {
	err := os.RemoveAll(getExecutionDataVolume())
	if err != nil {
		return err
	}
	err = os.RemoveAll(getConsensusDataVolume())
	if err != nil {
		return err
	}
	err = os.RemoveAll(getValidatorDataVolume())
	if err != nil {
		return err
	}
	return nil
}
