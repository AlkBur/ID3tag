package ID3tag

import (
	"io"
	"golang.org/x/text/encoding"
	"log"
	"golang.org/x/text/encoding/unicode"
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

func getLongAt(r io.ReaderAt, iOffset int64, bBigEndian bool) (int, error) {
	b, err := getBytesAt(r, iOffset, 4)
	if err != nil  {
		return -1, err
	}
	var n int
	if bBigEndian {
		for _, x := range b {
			n = n << 8
			n |= int(x)
		}
	}else{
		for i:=4; i>0; i-- {
			x := b[i-1]
			n = n << 8
			n |= int(x)
		}
	}
	if (n < 0) {
		n += 4294967296
	}
	return n, nil
}

func getStringAt(r io.ReaderAt, offset int64, n int) (string, error) {
	b, err := getBytesAt(r, offset, n)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getInteger24At(r io.ReaderAt, iOffset int64, bBigEndian bool) (int, error) {
	b, err := getBytesAt(r, iOffset, 3)
	if err != nil {
		return 0, err
	}
	var iInteger int
	if bBigEndian {
		for _, x := range b {
			iInteger = iInteger << 8
			iInteger |= int(x)
		}
	} else {
		for i := 4; i > 0; i-- {
			x := b[i-1]
			iInteger = iInteger << 8
			iInteger |= int(x)
		}
	}

	if (iInteger < 0) {
		iInteger += 16777216
	}
	return iInteger, err
}

func readSynchsafeInteger32At(r io.ReaderAt, iOffset int64) (int, error) {
	b, err := getBytesAt(r, iOffset, 4)
	if err != nil {
		return -1, err
	}
	var size int
	for _, x := range b {
		size = size << 7
		size |= int(x)
	}

	return size, err
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


	//var o binary.ByteOrder
	////ix := 2
	////if( b[0] == 0xFE && b[1] == 0xFF ) {
	//	o = binary.BigEndian
	////}else {
	////	o = binary.LittleEndian
	////}
	//
	//utf := make([]uint16, (len(b)+(2-1))/2)
	//for i := 0; i+(2-1) < len(b); i += 2 {
	//	utf[i/2] = o.Uint16(b[i:])
	//}
	//if len(b)/2 < len(utf) {
	//	utf[len(utf)-1] = utf8.RuneError
	//}
	//return string(utf16.Decode(utf))

}

func readUTF8String(b []byte) string {
	return string(b)
}

func readNullTerminatedString(b []byte) string {
	return string(b)
}