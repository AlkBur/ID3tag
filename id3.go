package ID3tag

import (
	"os"
	"errors"
	"fmt"
)

const (
	iso_8859_1 = iota
	utf_16
	utf_16le
	utf_16be
	utf_8
)

var (
	ErrorFormatID3 = errors.New("Error format file")
)

type ID3 struct {
	version string
	subVer  int
	size    int64
	f       *os.File
	Path string

	title   []byte //Название трека
	artist  []byte //Исполнитель
	album   []byte //Название альбома
	year    []byte    //Год
	comment []byte //Комментарий
	track   []byte    //Номер трека
	genre   []byte    //Индекс жанра
}

func New(f *os.File) (*ID3, error) {
	id3 := &ID3{f: f}
	err := id3.readTag()
	if err != nil {
		return nil, err
	}
	return id3, nil
}

func (id3 *ID3) readTag() error {
	if info, err := id3.f.Stat(); err != nil {
		return err
	}else{
		id3.size = info.Size()
	}
	id3.Path = id3.f.Name()
	if id3.size < 128 {
		return ErrorFormatID3
	}
	b, err := getBytesAt(id3.f, 0, 10)
	if err != nil {
		return err
	}
	marker := string(b[:3])
	if marker != "ID3" {
		b, err = getBytesAt(id3.f, id3.size-128, 128)
		if err != nil {
			return err
		}
		marker = string(b[:3])
		if marker != "TAG" {
			return ErrorFormatID3
		}
		id3.version = "ID3v1"
		err = id3.readID3v1(b)
	}else{
		id3.version = fmt.Sprintf("ID3v2.%d", b[3])
		id3.subVer = int(b[3])
		unsynch := isBitSetAt(b[5], 7)
		xheader := isBitSetAt(b[5], 6)
		xindicator := isBitSetAt(b[5], 5)
		size := getIntWithoutBit(b[6:])

		err = id3.readID3v2(size, unsynch, xheader, xindicator)
	}
	return err
}

func (id3 *ID3) Title() string {
	return decodeText(id3.title)
}

func (id3 *ID3) Comment() string {
	return decodeText(id3.comment)
}

func (id3 *ID3) Album() string {
	return decodeText(id3.album)
}

func (id3 *ID3) Artist() string {
	return decodeText(id3.artist)
}

func (id3 *ID3) Year() string {
	return decodeText(id3.year)
}

func ReadPath(path string) ([]*ID3, error) {
	ids := make([]*ID3, 0)

	f, err := os.Open(path)
	if err != nil {
		return ids, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return ids, err
	}else if !info.IsDir() {
		id, err := New(f)
		if err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}else{

	}
	return ids, nil
}