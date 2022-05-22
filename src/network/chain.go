package network

import "strings"

type Chain int64

const (
	MainNet Chain = iota
	L16
	L16Beta
	Local

	DefaultNetworkID = "l16beta"
)

/*
	Currently Supported Networks
*/
func IsChainSupported(chain Chain) bool {
	return chain == L16Beta || chain == Local
}

func (c Chain) String() string {
	switch c {
	case MainNet:
		return ChainMainNet
	case L16:
		return ChainL16
	case L16Beta:
		return ChainL16Beta
	case Local:
		return ChainLocal
	default:
		return "unknown chain"
	}
}

func GetChainByString(chainId string) Chain {
	c := strings.ToLower(chainId)
	switch c {
	case "l16":
		return L16
	case "l16beta":
		return L16Beta
	case "mainnet":
		return MainNet
	case "local":
		return Local
	default:
		return L16Beta
	}
}
