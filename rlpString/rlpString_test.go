package rlpString

import (
	"math/big"
	"testing"
)

func TestNewRlpString(t *testing.T) {
	var rightString = "0x48656c6c6f20576f726c6421"
	var value = "Hello World!"
	rString := CreateRlpString(value)

	if rString.AsString() != rightString {
		t.Errorf("CreateRlpString incorrect, got: %s, want: %s.", rString.AsString(), rightString)
	}

	rString = CreateRlpStringBigInt(big.NewInt(10))
	rightString = "0x0a"

	if rString.AsString() != rightString {
		t.Errorf("CreateRlpStringBigInt incorrect, got: %s, want: %s.", rString.AsString(), rightString)
	}

	rString = CreateRlpStringBigInt(big.NewInt(-10))
	if rString.AsString() != rightString {
		t.Errorf("CreateRlpStringBigInt incorrect, got: %s, want: %s.", rString.AsString(), "0x0a")
	}

}
