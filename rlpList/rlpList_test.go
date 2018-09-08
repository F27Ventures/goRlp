package rlpList

import (
	"testing"

	rlpString "github.com/gorlp/rlpString"
)

func TestNewRlpList(t *testing.T) {
	var value = "Hello World!"
	var rightString = "0x48656c6c6f20576f726c6421"
	rString := rlpString.CreateRlpString(value)
	rList := NewRlpListVariadic(rString)

	r := rList.GetValue()[0].(*rlpString.RlpString)

	if r.AsString() != rightString {
		t.Errorf("CreateRlpString incorrect, got: %s, want: %s.", r.AsString(), rightString)
	}

}
