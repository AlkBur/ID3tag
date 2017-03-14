package ID3tag

import (
	"errors"
)

var ErrorVersin24 = errors.New("Error version. Version > 2.4")

var _frames = map[string]string{
	// v2.2
	"BUF": "Recommended buffer size",
	"CNT": "Play counter",
	"COM": "Comments",
	"CRA": "Audio encryption",
	"CRM": "Encrypted meta frame",
	"ETC": "Event timing codes",
	"EQU": "Equalization",
	"GEO": "General encapsulated object",
	"IPL": "Involved people list",
	"LNK": "Linked information",
	"MCI": "Music CD Identifier",
	"MLL": "MPEG location lookup table",
	"PIC": "Attached picture",
	"POP": "Popularimeter",
	"REV": "Reverb",
	"RVA": "Relative volume adjustment",
	"SLT": "Synchronized lyric/text",
	"STC": "Synced tempo codes",
	"TAL": "Album/Movie/Show title",
	"TBP": "BPM (Beats Per Minute)",
	"TCM": "Composer",
	"TCO": "Content type",
	"TCR": "Copyright message",
	"TDA": "Date",
	"TDY": "Playlist delay",
	"TEN": "Encoded by",
	"TFT": "File type",
	"TIM": "Time",
	"TKE": "Initial key",
	"TLA": "Language(s)",
	"TLE": "Length",
	"TMT": "Media type",
	"TOA": "Original artist(s)/performer(s)",
	"TOF": "Original filename",
	"TOL": "Original Lyricist(s)/text writer(s)",
	"TOR": "Original release year",
	"TOT": "Original album/Movie/Show title",
	"TP1": "Lead artist(s)/Lead performer(s)/Soloist(s)/Performing group",
	"TP2": "Band/Orchestra/Accompaniment",
	"TP3": "Conductor/Performer refinement",
	"TP4": "Interpreted, remixed, or otherwise modified by",
	"TPA": "Part of a set",
	"TPB": "Publisher",
	"TRC": "ISRC (International Standard Recording Code)",
	"TRD": "Recording dates",
	"TRK": "Track number/Position in set",
	"TSI": "Size",
	"TSS": "Software/hardware and settings used for encoding",
	"TT1": "Content group description",
	"TT2": "Title/Songname/Content description",
	"TT3": "Subtitle/Description refinement",
	"TXT": "Lyricist/text writer",
	"TXX": "User defined text information frame",
	"TYE": "Year",
	"UFI": "Unique file identifier",
	"ULT": "Unsychronized lyric/text transcription",
	"WAF": "Official audio file webpage",
	"WAR": "Official artist/performer webpage",
	"WAS": "Official audio source webpage",
	"WCM": "Commercial information",
	"WCP": "Copyright/Legal information",
	"WPB": "Publishers official webpage",
	"WXX": "User defined URL link frame",
	// v2.3
	"AENC": "Audio encryption",
	"APIC": "Attached picture",
	"COMM": "Comments",
	"COMR": "Commercial frame",
	"ENCR": "Encryption method registration",
	"EQUA": "Equalization",
	"ETCO": "Event timing codes",
	"GEOB": "General encapsulated object",
	"GRID": "Group identification registration",
	"IPLS": "Involved people list",
	"LINK": "Linked information",
	"MCDI": "Music CD identifier",
	"MLLT": "MPEG location lookup table",
	"OWNE": "Ownership frame",
	"PRIV": "Private frame",
	"PCNT": "Play counter",
	"POPM": "Popularimeter",
	"POSS": "Position synchronisation frame",
	"RBUF": "Recommended buffer size",
	"RVAD": "Relative volume adjustment",
	"RVRB": "Reverb",
	"SYLT": "Synchronized lyric/text",
	"SYTC": "Synchronized tempo codes",
	"TALB": "Album/Movie/Show title",
	"TBPM": "BPM (beats per minute)",
	"TCOM": "Composer",
	"TCON": "Content type",
	"TCOP": "Copyright message",
	"TDAT": "Date",
	"TDLY": "Playlist delay",
	"TENC": "Encoded by",
	"TEXT": "Lyricist/Text writer",
	"TFLT": "File type",
	"TIME": "Time",
	"TIT1": "Content group description",
	"TIT2": "Title/songname/content description",
	"TIT3": "Subtitle/Description refinement",
	"TKEY": "Initial key",
	"TLAN": "Language(s)",
	"TLEN": "Length",
	"TMED": "Media type",
	"TOAL": "Original album/movie/show title",
	"TOFN": "Original filename",
	"TOLY": "Original lyricist(s)/text writer(s)",
	"TOPE": "Original artist(s)/performer(s)",
	"TORY": "Original release year",
	"TOWN": "File owner/licensee",
	"TPE1": "Lead performer(s)/Soloist(s)",
	"TPE2": "Band/orchestra/accompaniment",
	"TPE3": "Conductor/performer refinement",
	"TPE4": "Interpreted, remixed, or otherwise modified by",
	"TPOS": "Part of a set",
	"TPUB": "Publisher",
	"TRCK": "Track number/Position in set",
	"TRDA": "Recording dates",
	"TRSN": "Internet radio station name",
	"TRSO": "Internet radio station owner",
	"TSIZ": "Size",
	"TSRC": "ISRC (international standard recording code)",
	"TSSE": "Software/Hardware and settings used for encoding",
	"TYER": "Year",
	"TXXX": "User defined text information frame",
	"UFID": "Unique file identifier",
	"USER": "Terms of use",
	"USLT": "Unsychronized lyric/text transcription",
	"WCOM": "Commercial information",
	"WCOP": "Copyright/Legal information",
	"WOAF": "Official audio file webpage",
	"WOAR": "Official artist/performer webpage",
	"WOAS": "Official audio source webpage",
	"WORS": "Official internet radio station homepage",
	"WPAY": "Payment",
	"WPUB": "Publishers official webpage",
	"WXXX": "User defined URL link frame",
}

var _shortcuts = map[string]string{
	//Title
	"TIT2":   "Title",
	"TT2":  "Title",
	//Artist
	"TPE1": "Artist",
	"TP1": "Artist",
	//Album
	"TALB": "Album",
	"TAL": "Album",
	//Year
	"TYER": "Year",
	"TYE": "Year",
	//Comment
	"COMM": "Comment",
	"COM":"Comment",
	//Track
	"TRCK":"Track",
	"TRK":"Track",
	//Genre
	"TCON":"Genre",
	"TCO":"Genre",
	//Picture
	//"APIC": "Picture",
	//"PIC": "Picture",
	//Lyrics
	//"USLT":"Lyrics",
	//"ULT":"Lyrics",
}


func (id3 *ID3) readID3v2(size int, unsynch, xheader, xindicator bool) error {
	if int64(size) > id3.size {
		return ErrorFormatID3
	}
	if unsynch || xheader || xindicator {
		panic("Test unsynch || xheader || xindicator")
	}

	b := make([]byte, 10)
	offset := int64(10)
	for offset < int64(size) {
		frameHeaderSize  := 10
		if id3.subVer == 2 {
			frameHeaderSize  = 6
		}
		n, err := id3.f.ReadAt(b, offset)
		if n < frameHeaderSize  || err != nil {
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
		offsetflags := frameHeaderSize-2
		//grouping_identity := isBitSetAt(b[offsetflags+1], 7)
		//compression := isBitSetAt(b[offsetflags+1], 3)
		//encription := isBitSetAt(b[offsetflags+1], 2)
		unsynchronisation := isBitSetAt(b[offsetflags+1], 1)
		data_length_indicator := isBitSetAt(b[offsetflags+1], 0)



		iStart := offset + int64(frameHeaderSize)
		offset += int64(frameHeaderSize)  + int64(frameSize)

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
			id3.title, err = getBytesAt(id3.f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Artist":
			id3.artist, err = getBytesAt(id3.f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Album":
			id3.album, err = getBytesAt(id3.f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Year":
			id3.year, err = getBytesAt(id3.f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Comment":
			id3.comment, err = getBytesAt(id3.f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Track":
			id3.track, err = getBytesAt(id3.f, iStart, frameSize)
			if err != nil {
				return err
			}
		case "Genre":
			id3.genre, err = getBytesAt(id3.f, iStart, frameSize)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
