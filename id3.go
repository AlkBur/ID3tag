package ID3tag

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	iso_8859_1 = iota
	utf_16
	//utf_16le
	utf_16be
	utf_8
)

const (
	maxGoroutines      = 100
	maxSizeOfSmallFile = 1024 * 32
)

var (
	ErrorFormatID3 = errors.New("Error format file")
)

type ID3 struct {
	hash    string
	version string
	subVer  int
	size    int64
	path    string

	title   []byte //Название трека
	artist  []byte //Исполнитель
	album   []byte //Название альбома
	year    []byte //Год
	comment []byte //Комментарий
	track   []byte //Номер трека
	genre   []byte //Индекс жанра
}

func New(f *os.File) (*ID3, error) {
	id3 := &ID3{path: f.Name()}
	err := id3.readTag(f)
	if err != nil {
		return nil, err
	}
	return id3, nil
}

func (id3 *ID3) readTag(f *os.File) error {
	if info, err := f.Stat(); err != nil {
		return err
	} else {
		id3.size = info.Size()
	}
	if id3.size < 128 {
		return ErrorFormatID3
	}
	b, err := getBytesAt(f, 0, 10)
	if err != nil {
		return err
	}
	marker := string(b[:3])
	if marker != "ID3" {
		b, err = getBytesAt(f, id3.size-128, 128)
		if err != nil {
			return err
		}
		marker = string(b[:3])
		if marker != "TAG" {
			return ErrorFormatID3
		}
		id3.version = "ID3v1"
		err = id3.readID3v1(b)
	} else {
		id3.version = fmt.Sprintf("ID3v2.%d", b[3])
		id3.subVer = int(b[3])
		unsynch := isBitSetAt(b[5], 7)
		xheader := isBitSetAt(b[5], 6)
		xindicator := isBitSetAt(b[5], 5)
		size := getIntWithoutBit(b[6:])

		err = id3.readID3v2(f, size, unsynch, xheader, xindicator)
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

func (id3 *ID3) Path() string {
	return id3.path
}

func (id3 *ID3) FileName() string {
	return filepath.Base(id3.path)
}

func (id3 *ID3) Size() int64 {
	return id3.size
}

func ReadPath(path string) ([]*ID3, error) {
	var ids []*ID3

	f, err := os.Open(path)
	if err != nil {
		return ids, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return ids, err
	} else if !info.IsDir() {
		id, err := New(f)
		if err != nil {
			return ids, err
		}
		ids = make([]*ID3, 1)
		ids[0] = id
	} else {
		ids = make([]*ID3, 0, 20)
		id3Chan := make(chan *ID3, maxGoroutines*2)
		done := make(chan struct{})

		go func() {
			for id3 := range id3Chan {
				if id3 != nil {
					ids = append(ids, id3)
				}
			}
			done <- struct{}{}
		}()

		waiter := &sync.WaitGroup{}
		filepath.Walk(path, search(id3Chan, waiter))
		waiter.Wait()
		close(id3Chan)
		<-done
	}
	return ids, nil
}

func search(id3 chan *ID3, waiter *sync.WaitGroup) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && info.Size() > 128 &&
			(info.Mode()&os.ModeType == 0) {
			//get tag ID3
			if info.Size() < maxSizeOfSmallFile ||
				runtime.NumGoroutine() > maxGoroutines {
				getID3(path, id3)
			} else {
				waiter.Add(1)
				go func() {
					getID3(path, id3)
					waiter.Done()
				}()
			}
		}
		return nil
	}
}

func getID3(filename string, id3Chan chan *ID3) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("File:", filename, "; Error:", err)
		return
	}
	defer file.Close()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	id3, err := New(file)
	if err != nil {
		log.Println("File:", filename, "; Error:", err)
		return
	}

	//hash := sha1.New()
	//if size, err := io.Copy(hash, file);
	//	size != info.Size() || err != nil {
	//	if err != nil {
	//		log.Println("error:", err)
	//	} else {
	//		log.Println("error: failed to read the whole file:", filename)
	//	}
	//	return
	//}
	//id3.hash = string(hash.Sum(nil))

	id3Chan <- id3
}
