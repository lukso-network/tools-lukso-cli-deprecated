package network

import "strings"

type Chain int64

const (
	MainNet Chain = iota
	L16
	L16Beta
)

func GetChainByString(chainId string) Chain {
	c := strings.ToLower(chainId)
	switch c {
	case "l16":
		return L16
	case "L16beta":
		return L16Beta
	case "mainnet":
		return MainNet
	default:
		return L16Beta
	}
}
