package api

import (
	"fmt"
	"strings"
)

type IP struct {
	rawIP string
}

func NewIP(ip string) IP {
	return IP{ip}
}

func (ip IP) HttpAddressFromIP() string {
	if strings.Contains(ip.rawIP, "http") {
		return ip.rawIP
	}

	return fmt.Sprintf("http://%v", ip.rawIP)
}

func (ip IP) RPCAddress() string {
	return fmt.Sprintf("%v:8545", ip.HttpAddressFromIP())
}
