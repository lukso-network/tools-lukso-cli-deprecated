package gethrpc

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

type HexString struct {
	value []byte
}

func NewHexString() *HexString {
	return new(HexString)
}

func (it *HexString) Equal(other *HexString) bool {
	return bytes.Compare(it.value, other.value) == 0
}

func (it *HexString) SetBytes(b []byte) *HexString {
	it.value = b
	return it
}

func (it *HexString) SetBigInt(i *big.Int) *HexString {
	return NewHexString().SetBytes(i.Bytes())
}

func (it *HexString) SetInt64(i int64) *HexString {
	b := make([]byte, 8)

	binary.BigEndian.PutUint64(b, uint64(i))

	// trim byte array
	pos := 0
	for k, v := range b {
		if v > 0 {
			pos = k
			break
		}
	}

	return it.SetBytes(b[pos:])
}

// Given a string in the form "0x1234..." SetString treats it as
// a hex encoded string and transforms it to a HexString
// The string can be with "0x" prefix, without prefix and uneven
// letter are corrected by left pad with zeros
func (it *HexString) SetString(s string) (*HexString, error) {
	temp := s

	if strings.Contains(s, "0x") {
		temp = s[2:]
	}

	// normalize if not even
	if len(temp)%2 != 0 {
		temp = "0" + temp
	}

	b, err := hex.DecodeString(temp)

	if err != nil {
		return nil, err
	}

	return &HexString{
		value: b,
	}, nil
}

func (it *HexString) FormatString(plain bool, trim bool, etherClientFormat bool) string {
	if plain {
		return hex.EncodeToString(it.value)
	}

	// Ethereum Nodes are representing 0 as "0x0"
	if len(it.value) == 0 {
		if etherClientFormat {
			return "0x0"
		} else {
			return "0x00"
		}
	}

	s := hex.EncodeToString(it.value)

	// Ethereum Nodes are representing 0 as "0x0"
	if s == "00" {
		if etherClientFormat {
			return "0x0"
		} else {
			return "0x00"
		}
	}

	if trim {
		pos := 0
		for k, v := range s {
			if v != '0' {
				pos = k
				break
			}
		}

		return "0x" + s[pos:]
	}

	return "0x" + s
}

func (it *HexString) UntrimmedString() string {
	return it.FormatString(false, false, true)
}

func (it *HexString) String() string {
	return it.FormatString(false, false, true)
}

func (it *HexString) Bytes() []byte {
	return it.value
}

func (it *HexString) Plain() string {
	return it.FormatString(true, false, false)
}

// used to display text, from ascii 7 on there are meaningful chars
func (it *HexString) Ascii() string {
	b := it.Bytes()
	rb := make([]byte, 0)
	for _, v := range b {
		if v > 7 {
			rb = append(rb, v)
		}
	}

	return string(rb)
}

func (it HexString) BigInt() *big.Int {
	return new(big.Int).SetBytes(it.value)
}

func (it *HexString) Int64(bigEndian bool) int64 {
	b := it.Bytes()

	if len(b) == 0 {
		return 0
	}

	if len(b) > 8 {
		temp := b[len(b)-8:]
		b = temp
	}
	// pad if necessary
	if len(b) < 8 {
		temp := make([]byte, 8)

		for i := 8 - len(b); i < 8; i++ {
			temp[i] = b[i+len(b)-8]
		}

		b = temp
	}

	if bigEndian {
		return int64(binary.BigEndian.Uint64(b))
	}
	return int64(binary.LittleEndian.Uint64(b))
}

func (it *HexString) Concat(other *HexString) *HexString {
	nb := make([]byte, len(it.value)+len(other.value))

	copy(nb[:len(it.value)], it.value)
	copy(nb[len(it.value):], other.value)

	return new(HexString).SetBytes(nb)
}

func (it *HexString) LeftPadTo(length int) *HexString {
	if len(it.value) == length {
		return new(HexString).SetBytes(it.value)
	}

	nb := make([]byte, length)
	paddingLength := length - len(it.value)

	if len(it.value) > length {
		copy(nb, it.value[-paddingLength:])
	} else {
		copy(nb[paddingLength:], it.value)
	}

	return new(HexString).SetBytes(nb)
}

func (it *HexString) RightPadTo(length int) *HexString {
	if len(it.value) == length {
		return new(HexString).SetBytes(it.value)
	}

	nb := make([]byte, length)
	paddingLength := length - len(it.value)

	if len(it.value) > length {
		copy(nb, it.value[-paddingLength:])
	} else {
		copy(nb[paddingLength:], it.value)
	}

	return new(HexString).SetBytes(nb)
}

func (it *HexString) TrimLeft() *HexString {
	b1 := it.Bytes()

	for k, v := range b1 {
		if v != 0 {
			return new(HexString).SetBytes(b1[k:])
		}
	}

	return new(HexString).Empty()
}

func (it *HexString) TrimRight() *HexString {
	b1 := it.Bytes()

	for i := len(b1) - 1; i >= 0; i-- {
		if b1[i] != 0 {
			return new(HexString).SetBytes(b1[:i])
		}
	}

	return new(HexString).Empty()
}

// return an empty HexString with 0 bytes
func (it *HexString) Empty() *HexString {
	return &HexString{
		value: make([]byte, 0),
	}
}
func (it *HexString) Trimmed() string {
	return it.FormatString(false, true, true)
}

// List functions

// compares 2 hex string list
func CompareHexStringList(s1 []HexString, s2 []HexString) error {
	if len(s1) != len(s2) {
		return fmt.Errorf("wrong sizes in HexString List %v - %v", len(s1), len(s2))
	}

	for k, v := range s1 {
		if !v.Equal(&s2[k]) {
			return fmt.Errorf("not equal at position %v, %v : %v", k, v, s2[k])
		}
	}

	return nil
}

// transforms a string list with hex values to a hex string list
func StringListToHexStringList(l []string) ([]HexString, error) {
	result := make([]HexString, len(l))

	for k, v := range l {
		it, err := new(HexString).SetString(v)
		if err != nil {
			return nil, err
		}
		result[k] = *it
	}

	return result, nil
}

// transforms a hex string list into an untrimmed string list
func HexStringListToStringList(s []HexString) []string {
	itl := make([]string, len(s))

	for k, v := range s {
		itl[k] = v.UntrimmedString()
	}

	return itl
}
