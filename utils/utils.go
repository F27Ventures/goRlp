package utils

import (
	"fmt"
	"strings"
)

func ToHexStringByte(input []byte) string {
	return toHexString(input, 0, len(input), true)

}

func toHexString(input []byte, offset int, length int, withPrefix bool) string {
	var stringBuilder strings.Builder

	if withPrefix {
		stringBuilder.WriteString("0x")
	}

	for i := offset; i < offset+length; i++ {
		stringBuilder.WriteString(fmt.Sprintf("%02x", input[i]&0xff))
	}
	return stringBuilder.String()
}
