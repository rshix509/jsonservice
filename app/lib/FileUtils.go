package Utils

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
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

func (f FileInMem) ReadContentsByPart() (*io.PipeReader, string) {

	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	ct := writer.FormDataContentType()

	go func() {

		file, err := os.Open(f.Filename)
		if err != nil {
			log.Println("ERROR reading file" + f.Filename + " error: " + err.Error())
			pw.CloseWithError(err)
			return
		}
		defer file.Close()
		FormFile, err := writer.CreateFormFile("file", "large-file.json")
		if err != nil {
			log.Println("ERROR creating CreateFormFile" + f.Filename + " error: " + err.Error())
			pw.CloseWithError(err)
			return
		}
		_, err = io.Copy(FormFile, file)
		if err != nil {
			log.Println("ERROR copying" + f.Filename + " error: " + err.Error())
			pw.CloseWithError(err)
			return
		}
		pw.CloseWithError(writer.Close())
	}()

	return pr, ct

}
