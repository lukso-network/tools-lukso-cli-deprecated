package network

import "strings"

type Chain int64

const (
	MainNet Chain = iota
	L16
	Local
	Dev

	DefaultNetworkID = "mainnet"
)

/*
	Currently Supported Networks
*/
func IsChainSupported(chain Chain) bool {
	return chain == Dev || chain == Local || chain == L16
}

func (c Chain) String() string {
	switch c {
	case MainNet:
		return ChainMainNet
	case L16:
		return ChainL16
	case Local:
		return ChainLocal
	case Dev:
		return ChainDev
	default:
		return "unknown chain"
	}
}

func (c Chain) GetCurrencySymbol() string {
	switch c {
	case MainNet:
		return "LYX"
	default:
		return "LYXt"
	}
}

func GetChainByString(chainId string) Chain {
	c := strings.ToLower(chainId)
	switch c {
	case "l16":
		return L16
	case "mainnet":
		return MainNet
	case "local":
		return Local
	case "dev":
		return Dev
	default:
		return MainNet
	}
}
