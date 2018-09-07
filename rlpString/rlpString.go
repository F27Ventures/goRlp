package rlpString

import (
	"encoding/hex"
	"hash/fnv"
	"math/big"
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

func CreateRlpStringInt(value big.Int) *RlpString {
	if value.Cmp(big.NewInt(1)) < 0 {
		empty := []byte{}
		return NewRlpString(empty)
	} else {
		var bytes []byte = value.Bytes()
		if bytes[0] == 0 { // remove leading zero
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
	encodedStr := hex.EncodeToString(r.value)
	return encodedStr
}
func (r *RlpString) Hash() uint32 {
	str := r.AsString()
	return hash(str)
}

func (r *RlpString) GetBytes() []byte {
	return r.value
}
