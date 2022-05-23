package gethrpc

import "fmt"

func (client *Instance) GetBlock(number int64) (*Block, error) {

	hs := NewHexString().SetInt64(number)
	blockNumberInHex := hs.Trimmed()

	if number == 0 {
		blockNumberInHex = "0x0"
	}

	response, err := CheckRPCError(client.Call("eth_getBlockByNumber", blockNumberInHex, false))
	if err != nil {
		return nil, err
	}

	if response.Result == nil {
		return nil, fmt.Errorf("m: %v, p: %v didn't return error but also no response")
	}

	return toBlock(response.Result.(map[string]interface{}))
}

func toBlock(response map[string]interface{}) (*Block, error) {
	b := &Block{}
	b.Hash = response["hash"].(string)

	numberRaw := response["number"].(string)

	numberHS, err := NewHexString().SetString(numberRaw)
	if err != nil {
		return nil, err
	}
	b.Number = numberHS.Int64()

	transactionsRaw := response["transactions"].([]interface{})

	b.NumberOfTransactions = len(transactionsRaw)

	return b, nil
}

type Block struct {
	Hash                 string `json:"hash"`
	Number               int64  `json:"number"`
	NumberOfTransactions int    `json:"number_of_transactions"`
}
