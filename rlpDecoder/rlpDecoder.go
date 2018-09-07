package rlpDecoder

import (
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

func decode(rlpEncoded []byte) *rlpList.RlpList {
	rList := rlpList.NewRlpList(make([](interface{}), 0))
	Traverse(rlpEncoded, 0, len(rlpEncoded), rList)
	return rList
}

func Traverse(data []byte, startPos int, endPos int, rList *rlpList.RlpList) {
	if data == nil || len(data) == 0 {
		return
	}

	for startPos < endPos {
		var prefix int = int(data[startPos] & 0xff)

		if prefix < OFFSET_SHORT_STRING {
			// 1. the data is a string if the range of the
			// first byte(i.e. prefix) is [0x00, 0x7f],
			// and the string is the first byte itself exactly;
			rlpData := []byte{byte(prefix)}
			rList.SetValue(append(rList.GetValue(), rlpData.(interface{})...))
			startPos += 1
		} else if prefix == OFFSET_SHORT_STRING {
			// null
			emptyRstring = rlpString.NewRlpString([]byte{})
			startPos += 1

		} else if prefix > OFFSET_SHORT_STRING && prefix <= OFFSET_LONG_STRING {
			// 2. the data is a string if the range of the
			// first byte is [0x80, 0xb7], and the string
			// which length is equal to the first byte minus 0x80
			// follows the first byte;

			var strLen byte = byte(prefix - OFFSET_SHORT_STRING)

			var rlpData []byte = make([]byte, int(strLen))

			// copy(dst, src)
			copy(rlpData, data[startPos+1:startPos+1+int(strLen)])

			rList.SetValue(append(rList.GetValue(), rlpData.(interface{})...))

			startPos += 1 + strLen

		} else if prefix > OFFSET_LONG_STRING && prefix < OFFSET_SHORT_LIST {
			// 3. the data is a string if the range of the
			// first byte is [0xb8, 0xbf], and the length of the
			// string which length in bytes is equal to the
			// first byte minus 0xb7 follows the first byte,
			// and the string follows the length of the string

			var lenOfStrLen = byte(prefix - OFFSET_LONG_STRING)
			var strLen = calcLength(lenOfStrLen, data, startPos)

			rlpData := make([]byte, int(strLen))
			copy(rlpData, data[startPos+lenOfStrLen+1:startPos+lenOfStrLen+1+int(strLen)])

			rList.SetValue(append(rList.GetValue(), rlpData.(interface{})...))
			startPos += lenOfStrLen + strLen + 1

		} else if prefix >= OFFSET_SHORT_LIST && prefix <= OFFSET_LONG_LIST {
			// 4. the data is a list if the range of the
			// first byte is [0xc0, 0xf7], and the concatenation of
			// the RLP encodings of all items of the list which the
			// total payload is equal to the first byte minus 0xc0 follows the first byte;

			var listLen byte = byte(prefix - OFFSET_SHORT_LIST)

			newLevelList := rlpList.NewRlpList(make([](interface{}), 0))
			Traverse(rlpEncoded, startPos+1, startPos+listLen+1, newLevelList)

			rList.SetValue(append(rList.GetValue(), newLevelList.(interface{})...))
			startPos += 1 + listLen

		} else if prefix > OFFSET_LONG_LIST {
			// 5. the data is a list if the range of the
			// first byte is [0xf8, 0xff], and the total payload of the
			// list which length is equal to the
			// first byte minus 0xf7 follows the first byte,
			// and the concatenation of the RLP encodings of all items of
			// the list follows the total payload of the list;

			var lenOfListLen = byte(prefix - OFFSET_LONG_LIST)
			var listLen = calcLength(lenOfListLen, data, startPos)

			newLevelList := rlpList.NewRlpList(make([](interface{}), 0))
			Traverse(rlpEncoded, startPos+lenOfListLen+1, startPos+lenOfListLen+listLen+1, newLevelList)
			rList.SetValue(append(rList.GetValue(), newLevelList.(interface{})...))
			startPos += lenOfListLen + listLen + 1
		}
	}
}

func calcLength(lengthOfLength int, data []byte, pos int) int {
	var pow byte = byte(lengthOfLength - 1)
	var length int = 0

	for i := 0; i < lengthOfLength; i++ {
		length += int(data[pos+i]&0xff) << (8 * pow)
		pow--
	}
	return length
}
