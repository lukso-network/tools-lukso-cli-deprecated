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
	utils.ColoredPrintln("Is Finalized", participationResponse.Finalized)
	utils.ColoredPrintln("Participation Rate", participationResponse.Participation.GlobalParticipationRate)
	utils.ColoredPrintln("Eligible Ether", participationResponse.Participation.EligibleEther)
	utils.ColoredPrintln("Voted Ether", participationResponse.Participation.VotedEther)
	return nil
}
