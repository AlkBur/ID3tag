package ID3tag

import (
	"testing"
	"os"
)

const testFile = "tag.mp3"

func TestID3(t *testing.T)  {
	f,_ := os.Open(testFile)
	defer f.Close()

	tags,err := New(f)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Title", tags.Title())
	t.Log("Comment", tags.Comment())
	t.Log("Album", tags.Album())
	t.Log("Artist", tags.Artist())
	t.Log("Year", tags.Year())
}

//func TestINT7bit(t *testing.T)  {
//	b := []byte{0x00, 0x00, 0x07, 0x76}
//	t.Log(getIntWithoutBit(b))
//}
