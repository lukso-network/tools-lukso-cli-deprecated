package gethrpc

import "fmt"

/*
	Request Int64
*/
func (client *Instance) RequestInt64(method string, params ...interface{}) (int64, error) {
	response, err := CheckRPCError(client.Call(method, params...))
	if err != nil {
		return -1, err
	}

	if response.Result == nil {
		return -1, fmt.Errorf("m: %v, p: %v didn't return error but also no response", method, params)
	}

	val, ok := response.Result.(string)

	if !ok {
		return 0, fmt.Errorf("could not parse string from %s", response.Result)
	}

	hs, err := NewHexString().SetString(val)
	if err != nil {
		return -1, err
	}
	return hs.Int64(), nil
}
