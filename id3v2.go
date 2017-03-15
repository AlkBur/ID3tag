package ID3tag

import "io"

var _shortcuts = map[string]string{
	//Title
	"TIT2": "Title",
	"TT2":  "Title",
	//Artist
	"TPE1": "Artist",
	"TP1":  "Artist",
	//Album
	"TALB": "Album",
	"TAL":  "Album",
	//Year
	"TYER": "Year",
	"TYE":  "Year",
	//Comment
	"COMM": "Comment",
	"COM":  "Comment",
	//Track
	"TRCK": "Track",
	"TRK":  "Track",
	//Genre
	"TCON": "Genre",
	"TCO":  "Genre",
	//Picture
	//"APIC": "Picture",
	//"PIC": "Picture",
	//Lyrics
	//"USLT":"Lyrics",
	//"ULT":"Lyrics",
}

func (id3 *ID3) readID3v2(f io.ReaderAt, size int, unsynch, xheader, xindicator bool) error {
	if int64(size) > id3.size {
		return ErrorFormatID3
	}
	if unsynch || xheader || xindicator {
		panic("Test unsynch || xheader || xindicator")
	}

	b := make([]byte, 10)
	offset := int64(10)
	for offset < int64(size) {
		frameHeaderSize := 10
		if id3.subVer == 2 {
			frameHeaderSize = 6
		}
		n, err := f.ReadAt(b, offset)
		if n < frameHeaderSize || err != nil {
			return err
		}
		if b[0] == 0 {
			break
		}
		var ID string
		var frameSize int
		switch id3.subVer {
		case 2:
			ID = string(b[:3])
			frameSize = getInt(b[3:8])
		case 3:
			frameSize = getInt(b[4:8])
			ID = string(b[:4])
		case 4:
			frameSize = getIntWithoutBit(b[4:8])
			ID = string(b[:4])
		}
		offsetflags := frameHeaderSize - 2
		//grouping_identity := isBitSetAt(b[offsetflags+1], 7)
		//compression := isBitSetAt(b[offsetflags+1], 3)
		//encription := isBitSetAt(b[offsetflags+1], 2)
		unsynchronisation := isBitSetAt(b[offsetflags+1], 1)
		data_length_indicator := isBitSetAt(b[offsetflags+1], 0)

		iStart := offset + int64(frameHeaderSize)
		offset += int64(frameHeaderSize) + int64(frameSize)

		if data_length_indicator {
			panic("data length indicator")
			//frameDataSize = readSynchsafeInteger32At(frameDataOffset, frameData)
			//frameDataOffset += 4
			//frameSize -= 4
		}
		// TODO: support unsynchronisation
		if unsynchronisation {
			panic("unsynchronisation")
			continue
		}

		ids, ok := _shortcuts[ID]
		if !ok {
			continue
		}

		switch ids {
		case "Title":
			id3.title, err = getBytesAt(f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Artist":
			id3.artist, err = getBytesAt(f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Album":
			id3.album, err = getBytesAt(f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Year":
			id3.year, err = getBytesAt(f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Comment":
			id3.comment, err = getBytesAt(f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Track":
			id3.track, err = getBytesAt(f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Genre":
			id3.genre, err = getBytesAt(f, iStart, frameSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
