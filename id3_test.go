package ID3tag

import (
	"os"
	"testing"
)

const testFile = "v1tag.mp3"
const testPath = "Muzic"

func TestID3(t *testing.T) {
	t.Log("ID3")
	f, _ := os.Open(testFile)
	defer f.Close()

	tags, err := New(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Title", tags.Title())
	t.Log("Comment", tags.Comment())
	t.Log("Album", tags.Album())
	t.Log("Artist", tags.Artist())
	t.Log("Year", tags.Year())
}

func TestPath(t *testing.T) {
	t.Log("Path ID3")
	tags, err := ReadPath(testPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("tags", len(tags))
	//for _, tg := range tags {
	//	t.Log("Title", tg.Title())
	//	t.Log("Comment", tg.Comment())
	//	t.Log("Album", tg.Album())
	//	t.Log("Artist", tg.Artist())
	//	t.Log("Year", tg.Year())
	//	t.Log("=============================")
	//}
}

//func TestINT7bit(t *testing.T)  {
//	b := []byte{0x00, 0x00, 0x07, 0x76}
//	t.Log(getIntWithoutBit(b))
//}
