package reader

import (
	"encoding/json"
	"fmt"
	"os"
	"parser/internal/app/models"
)

type FileReader struct {
	filePath string
}

func NewFileReader(filePath string) *FileReader {
	return &FileReader{
		filePath: filePath,
	}
}

func (fileReader *FileReader) Read(dataCh chan<- models.User) {
	defer close(dataCh)

	file, err := os.Open(fileReader.filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	handleDelimiters(decoder)

	i := 1

	for decoder.More() {
		var user models.User

		if err := decoder.Decode(&user); err != nil {
			panic(fmt.Sprintf("%s (line: %d)", err, i))
		}

		dataCh <- user
	}

	handleDelimiters(decoder)
}

func handleDelimiters(decoder *json.Decoder) {
	if _, err := decoder.Token(); err != nil {
		panic(err)
	}
}
