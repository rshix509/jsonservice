package Utils

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/rshix509/jsonservice/app/lib/models"
)

var FileContent []models.EachEvent

type FileInMem struct {
	Filename string
}

func (f FileInMem) ReadContentsAndStoreStruct() {
	file, err := os.Open(f.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	buffer := ""
	var EachEvent models.EachEvent
	for scanner.Scan() {
		buffer = scanner.Text()
		buffer = buffer[1:]
		err = json.Unmarshal([]byte(buffer), &EachEvent)
		if err != nil {
			log.Print("WARN error decoding")
		}
		FileContent = append(FileContent, EachEvent)
	}
}

func (f FileInMem) ReadContentsByPart() *io.PipeReader {

	pr, pw := io.Pipe()

	go func() {

		file, err := os.Open(f.Filename)
		if err != nil {
			log.Println("ERROR reading file" + f.Filename + " error: " + err.Error())
			pw.CloseWithError(err)
			return
		}
		defer file.Close()

		_, err = io.Copy(pw, file)
		if err != nil {
			log.Println("ERROR copying" + f.Filename + " error: " + err.Error())
			pw.CloseWithError(err)
			return
		}
		pw.CloseWithError(err)
	}()

	return pr

}
