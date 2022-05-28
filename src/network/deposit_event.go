package network

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lukso-network/lukso-cli/api/gethrpc"
	"github.com/lukso-network/lukso-cli/src/network/contracts"
	"math/big"
)

type DepositEvents struct {
	Events []DepositEvent
}

type DepositEvent struct {
	PubKeyRaw             []byte
	Pubkey                string
	WithdrawalCredentials common.Hash
	Amount                *big.Int
	Signature             common.Hash
	Index                 int64
}

func NewDepositEvents(contractAddress, rpcEndpoint string) (DepositEvents, error) {
	events := DepositEvents{}
	client, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return events, err
	}
	contract, err := contracts.NewEth2Deposit(common.HexToAddress(contractAddress), client)
	if err != nil {
		return events, err
	}
	filterEvents, err := contract.FilterDepositEvent(nil)
	if err != nil {
		return events, err
	}
	events.Events = make([]DepositEvent, 0)
	for filterEvents.Next() {
		e, err := events.toEvent(filterEvents.Event)
		if err != nil {
			return events, err
		}
		events.Events = append(events.Events, e)
	}

	return events, nil
}

func (d DepositEvents) toEvent(fe *contracts.Eth2DepositDepositEvent) (DepositEvent, error) {
	e := DepositEvent{}

	e.PubKeyRaw = fe.Pubkey
	e.Pubkey = gethrpc.NewHexString().SetBytes(fe.Pubkey).String()
	e.WithdrawalCredentials = common.BytesToHash(fe.WithdrawalCredentials)
	e.Signature = common.BytesToHash(fe.Signature)

	e.Amount = big.NewInt(gethrpc.NewHexString().SetBytes(fe.Amount).Int64(false))
	e.Index = gethrpc.NewHexString().SetBytes(fe.Index).Int64(false)

	return e, nil
}

func (d DepositEvents) FindEventsWithKey(pubKey string) []DepositEvent {
	result := make([]DepositEvent, 0)
	for _, d := range d.Events {
		if pubKey == d.Pubkey {
			result = append(result, d)
		}
	}
	return result
}

func (d DepositEvents) Amount(pubKey string) int64 {
	events := d.FindEventsWithKey(pubKey)
	if len(events) == 0 {
		return 0
	}
	amount := int64(0)
	for _, e := range events {
		amount += e.Amount.Int64()
	}

	return amount
}
