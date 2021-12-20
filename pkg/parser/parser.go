package parser

import (
	"os"
)

func Parse(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		return ""
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return ""
	}

	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)

	if err != nil {
		return ""
	}

	return string(bs)
}
