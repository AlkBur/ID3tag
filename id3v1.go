package ID3tag

import "io"

type id3v1Reader struct {
	at io.ReaderAt
}

func (id3 *ID3) readID3v1(b []byte) error {
	if len(b)!=128 {
		return ErrorFormatID3
	}
	id3.title = b[3:33]
	id3.artist = b[33:63]
	id3.album = b[63:93]
	id3.year = b[93:97]
	id3.comment = b[97:127]
	//id3.Track = int(b[126])
	id3.genre = []byte{b[127]}

	return nil
}
