package ID3tag

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io"
	"log"
)

func getBytesAt(r io.ReaderAt, offset int64, size int) ([]byte, error) {
	b := make([]byte, size)
	n, err := r.ReadAt(b, offset)
	if n < size {
		err = ErrorFormatID3
	}
	return b, err
}

func isBitSetAt(b byte, iBit uint8) bool {
	return (b & (1 << iBit)) != 0
}

func getIntWithoutBit(b []byte) int {
	if len(b) != 4 {
		panic("Error get INT")
	}
	var size int
	for _, x := range b {
		size = size << 7
		size |= int(x)
	}
	return size
}

func getInt(b []byte) int {
	var size int
	for _, x := range b {
		size = size << 8
		size |= int(x)
	}
	return size
}

func getTextEncoding(b byte) int {
	var charset = -1
	switch b {
	case 0x00:
		charset = iso_8859_1
	case 0x01:
		charset = utf_16
	case 0x02:
		charset = utf_16be
	case 0x03:
		charset = utf_8
	}
	return charset
}

func decodeText(b []byte) string {
	if len(b) > 1 {
		switch getTextEncoding(b[0]) {
		case utf_16:
			return readUTF16String(b[1:], unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM))
		case utf_16be:
			return readUTF16String(b[1:], unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM))
		case utf_8:
			return readUTF8String(b[1:])
		default:
			return readUTF8String(b)
		}
	}
	return ""
}

func readUTF16String(b []byte, enc encoding.Encoding) string {
	decoder := enc.NewDecoder()
	dst := make([]byte, len(b)*2)
	nDst, _, err := decoder.Transform(dst, b, true)
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(dst[:nDst])
}

func readUTF8String(b []byte) string {
	return string(b)
}
