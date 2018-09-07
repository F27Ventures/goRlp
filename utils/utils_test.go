package utils

import (
	"testing"
)

func TestToHexString(t *testing.T) {

	var rightString = "0xe04fd020ea3a6910a2d808002b30309d"

	var test = []byte{byte(0xe0), byte(0x4f), byte(0xd0),
		byte(0x20), byte(0xea), byte(0x3a), byte(0x69), byte(0x10),
		byte(0xa2), byte(0xd8), byte(0x08), byte(0x00), byte(0x2b),
		byte(0x30), byte(0x30), byte(0x9d)}
	testString := ToHexStringByte(test)

	if testString != rightString {
		t.Errorf("ToHexString incorrect, got: %s, want: %s.",
			testString,
			rightString)
	}

}
