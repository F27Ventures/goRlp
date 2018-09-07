package rlpString

import (
	"testing"
)

func TestNewRlpString(t *testing.T) {
	var rightString = "0x48656c6c6f20576f726c6421"
	var value = "Hello World!"
	rString := CreateRlpString(value)

	if rString.AsString() != rightString {
		t.Errorf("Timesmape incorrect, got: %s, want: %s.", rString.AsString(), rightString)
	}

}
