package rlpEncoder

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	rlpString "github.com/gorlp/rlpString"
)

func TestNewRlpEncoder(t *testing.T) {
	//  83 as offset is 80 and length of string is 3 thus, 83
	// the d is 64, o is 6f, g is 67
	rightAnswer := "83646f67"
	rString := rlpString.CreateRlpString("dog")

	result := EncodeAll(rString)
	var buffer bytes.Buffer

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll dog incorrect, got: %s, want: %s.", buffer.String(), rightAnswer)
	}

	rightAnswer = "01"
	rString = rlpString.NewRlpString([]byte{byte(0x01)})

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll one byte incorrect, got: %s, want: %s.", buffer.String(), rightAnswer)
	}

	rightAnswer = "80"
	rString = rlpString.CreateRlpString("")

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}
	// fmt.Printf(buffer.String())

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll empytp string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}

	rightAnswer = "80"
	rString = rlpString.CreateRlpStringBigInt(big.NewInt(0))

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll 0 string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}

	rightAnswer = "820400"
	rString = rlpString.CreateRlpStringBigInt(big.NewInt(1024))

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll 0 string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}

	fmt.Printf(buffer.String())

}
