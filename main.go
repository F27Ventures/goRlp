package main

import "fmt"

/** RLP encoding rules are defined as follows: */

/*
 * For a single byte whose value is in the [0x00, 0x7f] range, that byte is
 * its own RLP encoding.
 */

/**
 * [0x80]
 * If a string is 0-55 bytes long, the RLP encoding consists of a single
 * byte with value 0x80 plus the length of the string followed by the
 * string. The range of the first byte is thus [0x80, 0xb7].
 */

var OFFSET_SHORT_ITEM int = 0x80

/**
 * [0xb7]
 * If a string is more than 55 bytes long, the RLP encoding consists of a
 * single byte with value 0xb7 plus the length of the length of the string
 * in binary form, followed by the length of the string, followed by the
 * string. For example, a length-1024 string would be encoded as
 * \xb9\x04\x00 followed by the string. The range of the first byte is thus
 * [0xb8, 0xbf].
 */

var OFFSET_LONG_ITEM int = 0xb7

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

/* ******************************************************
 * 						DECODING						*
 * ******************************************************/
func decodeOneByteItem(data []byte, index int) byte {
	if (int(data[index] & 0xFF)) == OFFSET_SHORT_ITEM {
		return (byte)(int(data[index]) - OFFSET_SHORT_ITEM)
	}

	if (int(data[index] & 0xFF)) < OFFSET_SHORT_ITEM {
		return (byte)(data[index])
	}

	if (int(data[index] & 0xFF)) == OFFSET_SHORT_ITEM+1 {
		return (byte)(data[index+1])
	}
	return 0
}

func decodeInt(data []byte, index int) int {
	var value int = 0
	if (int(data[index]&0xFF)) == OFFSET_SHORT_ITEM && (int(data[index]&0xFF)) < OFFSET_LONG_ITEM {
		var length byte = (byte)(int(data[index]) - OFFSET_SHORT_ITEM)
		var pow byte = (byte)(int(length) - 1)
		for i := 1; i <= int(length); i++ {
			value += int(data[index+i] << (8 * pow))
			pow--
		}

	} else {
		panic("wrong decode attempt")
	}
	return value
}

func plus(a int, b int) int {
	return a + b
}


func decode(data byte[], pos int) {
	if (data == nil || len(data) < 1) {
		return nil;
	}
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}
