package network

func Clear(configs *NodeConfigs) error {
	err := removeContents(configs.Execution.DataVolume)
	if err != nil {
		return err
	}
	err = removeContents(configs.Consensus.DataVolume)
	if err != nil {
		return err
	}
	err = removeContents(configs.Validator.DataVolume)
	if err != nil {
		return err
	}
	err = removeContents(configs.Configs.Volume)
	if err != nil {
		return err
	}
	return nil
}
