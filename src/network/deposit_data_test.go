package network

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseDepositDataFromFile(t *testing.T) {
	data, err := ParseDepositDataFromFile("../../assets/deposit_data.json")
	for _, myData := range data {
		rawData, _ := json.Marshal(myData)
		//t.Log(fmt.Sprintf("%+v", *myData))
		t.Log(string(rawData))
	}
	require.NoError(t, err)
}
