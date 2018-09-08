package rlpString

import (
	"hash/fnv"
	"math/big"

	utils "github.com/gorlp/utils"
)

type RlpString struct {
	value []byte
}

func NewRlpString(value []byte) *RlpString {
	return &RlpString{value: value}
}

func CreateRlpString(value string) *RlpString {
	return NewRlpString([]byte(value))
}

func CreateRlpStringBigInt(value *big.Int) *RlpString {
	if value.Sign() < 0 {
		return NewRlpString([]byte{})
	} else {
		var bytes = value.Bytes()
		if int(bytes[0]) == 0 { // remove leading zero
			return NewRlpString(bytes[1:])
		} else {
			return NewRlpString(bytes)
		}
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func (r *RlpString) AsString() string {
	encodedStr := utils.ToHexStringByte(r.value)
	return encodedStr
}
func (r *RlpString) Hash() uint32 {
	str := r.AsString()
	return hash(str)
}

func (r *RlpString) GetBytes() []byte {
	return r.value
}
