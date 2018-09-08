package rlpEncoder

import (
	"fmt"
	"reflect"
	"strings"

	rlpList "github.com/gorlp/rlpList"
	rlpString "github.com/gorlp/rlpString"
)

/**
 * [0x80]
 * If a string is 0-55 bytes long, the RLP encoding consists of a single
 * byte with value 0x80 plus the length of the string followed by the
 * string. The range of the first byte is thus [0x80, 0xb7].
 */

var OFFSET_SHORT_STRING int = 0x80

/**
 * [0xb7]
 * If a string is more than 55 bytes long, the RLP encoding consists of a
 * single byte with value 0xb7 plus the length of the length of the string
 * in binary form, followed by the length of the string, followed by the
 * string. For example, a length-1024 string would be encoded as
 * \xb9\x04\x00 followed by the string. The range of the first byte is thus
 * [0xb8, 0xbf].
 */
var OFFSET_LONG_STRING int = 0xb7

/**
 * [0xc0]
 * If the total payload of a list (i.e. the combined length of all its
 * items) is 0-55 bytes long, the RLP encoding consists of a single byte
 * with value 0xc0 plus the length of the list followed by the concatenation
 * of the RLP encodings of the items. The range of the first byte is thus
 * [0xc0, 0xf7].
 */
var OFFSET_SHORT_LIST int = 0xc0

/**
 * [0xf7]
 * If the total payload of a list is more than 55 bytes long, the RLP
 * encoding consists of a single byte with value 0xf7 plus the length of the
 * length of the list in binary form, followed by the length of the list,
 * followed by the concatenation of the RLP encodings of the items. The
 * range of the first byte is thus [0xf8, 0xff].
 */
var OFFSET_LONG_LIST int = 0xf7

func EncodeAll(value interface{}) []byte {
	varType := reflect.TypeOf(value).String()
	if strings.Contains(varType, "RlpString") {
		// fmt.Printf("instanisinationg RlpString ")
		return EncodeString(value.(*rlpString.RlpString))
	} else {
		return EncodeList(value.(*rlpList.RlpList))
	}
}

func EncodeString(value *rlpString.RlpString) []byte {
	return Encode(value.GetBytes(), OFFSET_SHORT_STRING)
}

func EncodeList(value *rlpList.RlpList) []byte {
	rlpTypes := value.GetValue()

	if len(rlpTypes) == 0 {
		return Encode([]byte{}, OFFSET_SHORT_LIST)
	} else {
		result := []byte{}
		for _, entry := range rlpTypes {
			result = append(result, EncodeAll(entry)...)
		}

		return Encode(result, OFFSET_SHORT_LIST)
	}
}

func Encode(bytesValue []byte, offset int) []byte {
	if len(bytesValue) == 1 &&
		offset == OFFSET_SHORT_STRING &&
		bytesValue[0] >= byte(0x00) &&
		bytesValue[0] <= byte(0x7f) {
		return bytesValue
	} else if len(bytesValue) < 55 {
		result := make([]byte, len(bytesValue)+1)
		result[0] = byte((offset + len(bytesValue)))
		copy(result[1:], bytesValue)
		return result
	} else {
		// TODO
		encodedStringLength := ToMinimalByteArray(len(bytesValue))
		result := make([]byte, len(bytesValue)+len(encodedStringLength)+1)
		result[0] = byte((offset + 0x37) + len(encodedStringLength))
		copy(result[1:len(encodedStringLength)+1], encodedStringLength)
		copy(result[len(encodedStringLength)+1:], bytesValue)
		return result
	}
}

func ToMinimalByteArray(value int) []byte {
	encoded := ToByteArry(value)

	for i := 0; i < len(encoded); i++ {
		if int(encoded[i]) != 0 {
			result := make([]byte, len(encoded)-i)
			copy(result, encoded[i:])
			return result
		}
	}

	return []byte{}
}

func ToByteArry(value int) []byte {
	result := []byte{byte(((value >> 24) & 0xff)),
		byte(((value >> 16) & 0xff)),
		byte(((value >> 8) & 0xff)),
		byte(((value) & 0xff))}
	return result
}

func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
