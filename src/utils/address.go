package utils

import (
	"fmt"
	"strings"
)

func MaybeAddHexPrefix(address string) string {
	a := address
	if !strings.Contains(address, "0x") {
		a = fmt.Sprintf("0x%s", address)
	}
	return strings.ToLower(a)
}
