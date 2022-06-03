package network

import "os"

func Clear(configs *NodeConfigs) error {
	err := os.RemoveAll(configs.Execution.DataVolume)
	if err != nil {
		return err
	}
	err = os.RemoveAll(configs.Consensus.DataVolume)
	if err != nil {
		return err
	}
	err = os.RemoveAll(configs.Validator.DataVolume)
	if err != nil {
		return err
	}
	err = os.RemoveAll(configs.Configs.Volume)
	if err != nil {
		return err
	}
	return nil
}
