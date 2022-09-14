package network

import (
	"github.com/lukso-network/lukso-cli/api/beaconapi"
	"github.com/lukso-network/lukso-cli/src/utils"
	"strconv"
)

func DescribeNetwork(baseUrl string, epoch int64) error {
	client := beaconapi.NewBeaconClient(baseUrl)
	chainHeadResponse, err := client.ChainHead()
	if err != nil {
		return err
	}

	utils.ColoredPrintln("Head Epoch", chainHeadResponse.HeadEpoch)
	utils.ColoredPrintln("Finalized Epoch", chainHeadResponse.FinalizedEpoch)
	utils.ColoredPrintln("Finalized Block Root", chainHeadResponse.FinalizedBlockRoot)
	utils.ColoredPrintln("Justified Epoch", chainHeadResponse.JustifiedEpoch)
	utils.ColoredPrintln("Justified Block Root", chainHeadResponse.PreviousJustifiedBlockRoot)

	headEpoch := epoch
	if epoch == -1 {
		headEpoch, err = strconv.ParseInt(chainHeadResponse.HeadEpoch, 10, 64)
		if err != nil {
			return err
		}
	}

	participationResponse, err := client.Participation(headEpoch)
	if err != nil {
		return err
	}

	utils.ColoredPrintln("Participation Epoch", participationResponse.Epoch)
	//TODO: is this needed? What does it mean.
	utils.ColoredPrintln("Is Finalized", participationResponse.Finalized)
	// TODO: In percentages
	utils.ColoredPrintln("Participation Rate", participationResponse.Participation.GlobalParticipationRate)
	//TODO: convert from Gwei to LYX
	utils.ColoredPrintln("Eligible LYX", participationResponse.Participation.EligibleEther)
	utils.ColoredPrintln("Voted LYX", participationResponse.Participation.VotedEther)
	return nil
}
